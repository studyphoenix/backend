package controller

import (
	"errors"
	"fmt"

	"github.com/TwiN/logr"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/studyphoenix/backend/httputil"
	"github.com/studyphoenix/backend/model"
)

// Request handler to get a user's flashcards
//
// @Summary 	Search for flashcards or get all [WIP]
// @Description	Returns a list of user flashcards, without any subtopics. Those need to be requested separately
// @Tags		flashcards
// @Produce		json
// @Param 		q query string false "The search query (matches name)"
// @Param 		count query uint false "The maximum number of results to return. Defaults to 25"
// @Param 		offset query uint false "The offset in the results list. Defaults to 0"
// @Param 		subtopicId query uint false "The subtopic the results should be associated with"
// @Param 		minCompletionCount query uint false "The minimum number of times the flashcard was completed"
// @Param 		maxCompletionCount query uint false "The maximum number of times the flashcard was completed"
// @Success		200	{object}		[]model.Flashcard
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/flashcards/search [get]
func (m *Controller) QueryFlashcards(c *gin.Context) {
	var req model.FlashcardQuery
	err := c.BindQuery(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse query: ", err))
		c.AbortWithStatus(400)
		return
	}

	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}
	flashcards, err := m.db.QueryFlashcards(req, uid)
	if err != nil {
		httputil.NewError(c, 500, errors.New("Failed to process query request"))
		c.Abort()
	} else {
		c.JSON(201, flashcards)
	}
}
