package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hci-group-15/backend/docs"
	"github.com/janishutz/oidclogin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddRoutes(c *Controller, r *gin.Engine) {
	// Assets
	assets := r.Group("/assets")
	assets.PUT("create", c.CreateAsset)
	assets.DELETE("delete/:id", c.DeleteAsset)
	assets.GET("get/:id", c.GetAssetDetails)
	assets.GET("download/:id", c.DownloadAsset)
	assets.PATCH("update/:id", c.UpdateAssetDeteails)
	assets.PUT("upload/:id", c.UploadAsset)

	// Chats
	chats := r.Group("/chats")
	chats.PUT("create")
	chats.DELETE("delete/:id", c.DeleteChat)
	chats.GET("get/:id", c.GetChat)
	chats.PUT("messages/add/:id", c.AddChatMessageToChat)
	chats.DELETE("messages/delete/:id", c.DeleteChatMessage)
	chats.PATCH("messages/update/:id", c.UpdateChatMessage)
	chats.PATCH("update/:id", c.UpdateChat)

	// Flashcards
	flashcards := r.Group("/flashcards")
	flashcards.PUT("create", c.CreateFlashcard)
	flashcards.DELETE("delete/:id", c.DeleteFlashcard)
	flashcards.GET("get/:id", c.GetFlashcard)
	flashcards.GET("search", c.QueryFlashcards)
	flashcards.PATCH("update/:id", c.UpdateFlashcard)
	flashcards.PATCH("complete/:id", c.CompleteFlashcard)

	// Subjects
	subjects := r.Group("/subjects")
	subjects.PATCH("associate/assets/:id", c.AssociateAssetToSubject)
	subjects.PATCH("associate/subtopic/:id", c.AssociateSubtopicToSubject)
	subjects.PUT("create", c.CreateSubject)
	subjects.PATCH("deassociate/assets/:id", c.DeassociateAssetToSubject)
	subjects.PATCH("deassociate/subtopic/:id", c.DeassociateSubtopicToSubject)
	subjects.DELETE("delete/:id", c.DeleteSubject)
	subjects.GET("get/:id", c.GetSubject)
	subjects.GET("search", c.QuerySubjects)
	subjects.GET("techtree/:id", c.GetSubjectTechTree)
	subjects.PATCH("update/:id", c.UpdateSubject)

	// Subtopics
	subtopics := r.Group("/subtopics")
	subtopics.PATCH("associate/assets/:id", c.AssociateAssetToSubtopic)
	subtopics.PATCH("associate/subtopic/:id", c.AssociateSubtopicToSubtopic)
	subtopics.PATCH("associate/flashcards/:id", c.AssociateFlashcardToSubtopic)
	subtopics.PATCH("associate/chats/:id", c.AssociateChatToSubtopic)
	subtopics.PUT("create", c.CreateSubtopic)
	subtopics.PATCH("deassociate/assets/:id", c.DeassociateAssetToSubtopic)
	subtopics.PATCH("deassociate/subtopic/:id", c.DeassociateSubtopicToSubtopic)
	subtopics.PATCH("deassociate/flashcards/:id", c.DeassociateFlashcardToSubtopic)
	subtopics.PATCH("deassociate/chats/:id", c.DeassociateChatToSubtopic)
	subtopics.DELETE("delete/:id", c.DeleteSubtopic)
	subtopics.GET("get/:id", c.GetSubtopic)
	subtopics.PATCH("update/:id", c.UpdateSubtopic)

	// Users
	users := r.Group("/user", oidclogin.EnsureLogin(false))
	users.GET("profile", c.GetUserProfile)

	// ── Healthcheck ──────────────────────────────────────────────────
	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"status": true}) })

	// ── Swagger ──────────────────────────────────────────────────────
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.BasePath = "/"
}
