package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/TwiN/logr"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hci-group-15/backend/httputil"
	"github.com/hci-group-15/backend/model"
)

// Request handler to create a new flashcard
//
// @Summary 	Create a new flashcard [WIP]
// @Description	Creates a new flashcard linked to the current user
// @Tags		flashcards
// @Produce		json
// @Param 		payload body model.CreateFlashcardRequest true "flashcard details to store"
// @Success		201	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/flashcards/create [put]
func (m *Controller) CreateFlashcard(c *gin.Context) {
	var req model.CreateFlashcardRequest
	err := c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create flashcard: ", err))
		c.AbortWithStatus(400)
		return
	}

	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}
	if m.db.CreateFlashcard(req, uid) {
		c.JSON(200, nil)
	} else {
		httputil.NewError(c, 500, errors.New("ERR_CREATE_FLASHCARD"))
	}
}

// Request handler to get a flashcard
//
// @Summary 	Get the specified flashcard [WIP]
// @Description	Returns the specified flashcard details, including extra metadata
// @Tags		flashcards
// @Produce		json
// @Param 		id path string true "The id of the flashcard to retrieve"
// @Success		200	{object}		model.Flashcard
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/flashcards/get/{id} [get]
func (m *Controller) GetFlashcard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	chat, err := m.db.GetFlashcard(uint(id))
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, chat)
}

// Request handler to update a flashcard
//
// @Summary 	Updates the specified flashcard [WIP]
// @Description	Updates the flashcard. Note that all attributes in the body are optional and if omitted, will not be updated
// @Tags		flashcards
// @Produce		json
// @Param 		id path string true "The id of the flashcard to update"
// @Param 		payload body model.UpdateFlashcardRequest true "Flashcard details to update with. All attributes optional!"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "flashcard not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/flashcards/update/{id} [patch]
func (m *Controller) UpdateFlashcard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.UpdateFlashcardRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to update flashcard, error: ", err))
		c.Abort()
		return
	}

	if m.db.UpdateFlashcard(uint(id), req) {
		c.JSON(200, nil)
	} else {
		c.JSON(404, nil)
	}
}

// Request handler to mark a flashcard as completed (again)
//
// @Summary 	Mark flashcard as completed (again) [WIP]
// @Description	Increments the completion count of the specified flashcard by one. Request body is ignored
// @Tags		flashcards
// @Produce		json
// @Param 		id path string true "The id of the flashcard to mark"
// @Success		204	{object}		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "flashcard not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/flashcards/complete/{id} [patch]
func (m *Controller) CompleteFlashcard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	if m.db.CompleteFlashcard(uint(id)) {
		c.JSON(200, nil)
	} else {
		c.JSON(404, nil)
	}
}

// Request handler to delete a flashcard
//
// @Summary 	Deletes the given flashcard [WIP]
// @Description	Deletes the given flashcard including all subtopics, assets and flashcards.
// @Tags		flashcards
// @Produce		json
// @Param 		id path string true "The id of the flashcard to update"
// @Success		204	{object}		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "flashcard not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/flashcards/delete/{id} [delete]
func (m *Controller) DeleteFlashcard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	if m.db.DeleteFlashcard(uint(id)) {
		c.JSON(200, nil)
	} else {
		c.JSON(404, nil)
	}
}
