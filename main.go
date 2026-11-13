package main

import (
	"os"
	"time"

	"github.com/TwiN/logr"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hci-group-15/backend/config"
	"github.com/hci-group-15/backend/controller"
	"github.com/hci-group-15/backend/model"
	"github.com/hci-group-15/backend/util"
	"github.com/janishutz/oidclogin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
)

// @title StudyPhoenix backend
// @version         0.1.0
// @description.markdown
// @termsOfService  https://janishutz.com/legal/tos/
//
// @contact.name   Support
// @contact.url    https://github.com/studyphoenix/backend#readme
// @contact.email  support@janishutz.com
//
// @license.name  GNU Affero General Public License 3.0
// @license.url   https://www.gnu.org/licenses/agpl-3.0.en.html
// @schemes 	  http https
//
// @tag.name assets
// @tag.description Asset CRUD (assets are any type of file the user can upload) [WIP]
//
// @tag.name chats
// @tag.description Chat CRUD, chat message CRUD [WIP]
//
// @tag.name flashcards
// @tag.description Flashcard CRUD [WIP]
//
// @tag.name subjects
// @tag.description Subjects CRUD, association of subtopics and assets [WIP]
//
// @tag.name subtopics
// @tag.description Subtopics CRUD, association of other subtopics, chats and assets [WIP]
//
// @tag.name user
// @tag.description User data retrieval and deletion
//
// @BasePath  /
func main() {
	logr.SetThreshold(logr.LevelDebug)
	logr.Debug("Starting...")
	r := gin.Default()

	conf := config.Load()
	prod := os.Getenv("PROD") == "true"

	// Load templates
	r.LoadHTMLGlob("templates/*.tmpl")

	// ┌                                               ┐
	// │            Security configuration             │
	// └                                               ┘
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{conf.Urls.FrontendURL},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(util.DefaultHeaders)

	if !conf.TestingMode {
		r.Use(util.RateLimiter())
	}
	// r.SetTrustedProxies(conf.Urls.TrustedProxies)

	// ┌                                               ┐
	// │                    Session                    │
	// └                                               ┘
	secret, ok := os.LookupEnv("SESSION_SECRET")
	if !ok {
		logr.Warn("Missing secret (SESSION_SECRET is unset), using unsafe default")
		secret = "secret"
	}
	store := memstore.NewStore([]byte(secret))
	r.Use(sessions.Sessions("session", store))

	// ┌                                               ┐
	// │               Routes, DB, Login               │
	// └                                               ┘
	db := model.Init(prod, &conf)
	oidclogin.ConfigureFull(r, conf.Urls.BackendURL, conf.Urls.DefaultRedirect, func(userid, name, email string, c *gin.Context) {
		uid := db.UpdateOrCreateUser(userid, name, email)
		ses := sessions.Default(c)
		ses.Set("userid", uid)
		ses.Set("username", name)
		ses.Set("email", email)
	}, true, nil, prod)

	// Routes
	c := controller.NewController(db)

	if conf.TestingMode {
		r.GET("/", oidclogin.EnsureLogin(false), func(c *gin.Context) {
			c.JSON(200, gin.H{"success": true})
		})
	} else {
		r.GET("/", func(c *gin.Context) {
			c.Redirect(307, "/swagger/index.html")
		})
	}

	c.AddRoutes(r)

	r.Run()
}
