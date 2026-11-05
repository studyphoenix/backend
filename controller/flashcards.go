package controller

import (
	"github.com/gin-gonic/gin"
)

type CreateFlashcardRequest struct {
	Name string `example:"Analysis" json:"name" maxLength:"250" minLength:"3"`
	Icon string `example:"fa-gear" json:"icon" maxLength:"250" minLength:"0"` // The Icon of the flashcard. Must be an icon that can be rendered in the browser
}

// Request handler to create a new flashcard
//
// @Summary 	Create a new flashcard [WIP]
// @Description	Creates a new flashcard linked to the current user
// @Tags		flashcards
// @Produce		json
// @Param 		payload body CreateFlashcardRequest true "flashcard details to store"
// @Success		201	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/flashcards/create [put]
func (m *Controller) CreateFlashcard(c *gin.Context) {
	c.JSON(200, nil)
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
	c.JSON(200, nil)
}

type UpdateFlashcardRequest struct {
	Name *string `example:"Analysis" json:"name" maxLength:"250" minLength:"3"`
	Icon *string `example:"fa-gear" json:"icon" maxLength:"250" minLength:"0"` // The Icon of the flashcard. Must be an icon that can be rendered in the browser
}

// Request handler to update a flashcard
//
// @Summary 	Updates the specified flashcard [WIP]
// @Description	Updates the flashcard. Note that all attributes in the body are optional and if omitted, will not be updated
// @Tags		flashcards
// @Produce		json
// @Param 		id path string true "The id of the flashcard to update"
// @Param 		payload body UpdateFlashcardRequest true "Flashcard details to update with. All attributes optional!"
// @Success		200	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "flashcard not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/flashcards/update/{id} [patch]
func (m *Controller) UpdateFlashcard(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to delete a flashcard
//
// @Summary 	Deletes the given flashcard [WIP]
// @Description	Deletes the given flashcard including all subtopics, assets and flashcards.
// @Tags		flashcards
// @Produce		json
// @Param 		id path string true "The id of the flashcard to update"
// @Success		200	{object}		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "flashcard not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/flashcards/delete/{id} [delete]
func (m *Controller) DeleteFlashcard(c *gin.Context) {
	c.JSON(200, nil)
}
