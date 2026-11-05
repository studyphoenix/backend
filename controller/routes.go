package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hci-group-15/backend/docs"
	"github.com/janishutz/oidclogin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddRoutes(c *Controller, r *gin.Engine) {
	// Users
	users := r.Group("/user", oidclogin.EnsureLogin(false))
	users.GET("profile", c.GetUserProfile)

	// Subjects
	subjects := r.Group("/subjects")
	subjects.GET("search", c.QuerySubjects)

	// Subtopics
	subtopics := r.Group("/subtopics")
	subtopics.GET("get/:id", c.GetSubtopic)

	// Flashcards
	flashcards := r.Group("/flashcards")
	flashcards.GET("search", c.QueryFlashcards)

	// Chats
	chats := r.Group("/chats")
	chats.GET("get/:id", c.GetChat)

	// Flashcards
	assets := r.Group("/assets")
	assets.GET("get/:id", c.QueryFlashcards)

	// ── Healthcheck ──────────────────────────────────────────────────
	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"status": true}) })

	// ── Swagger ──────────────────────────────────────────────────────
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.BasePath = "/"
}
