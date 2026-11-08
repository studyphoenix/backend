package controller

import (
	"github.com/gin-gonic/gin"
)

type CreateSubjectRequest struct {
	Name string `example:"Analysis" json:"name" maxLength:"100" minLength:"3" validate:"required,printascii,min=3,max=100"`
	Icon string `example:"fa-gear" json:"icon" maxLength:"100" minLength:"0" validate:"required,printascii,min=0,max=100"` // The Icon of the subject. Must be an icon that can be rendered in the browser
}

// Request handler to create a new subject
//
// @Summary 	Create a new subject [WIP]
// @Description	Creates a new subject linked to the current user
// @Tags		subjects
// @Produce		json
// @Param 		payload body CreateSubjectRequest true "Subject details to store"
// @Success		201	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/create [put]
func (m *Controller) CreateSubject(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to get a subject
//
// @Summary 	Get the specified subject [WIP]
// @Description	Returns the specified subject details, including extra metadata
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The id of the subject to retrieve"
// @Success		200	{object}		model.Subject
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/get/{id} [get]
func (m *Controller) GetSubject(c *gin.Context) {
	c.JSON(200, nil)
}

type UpdateSubjectRequest struct {
	Name string `example:"Analysis" json:"name" maxLength:"100" minLength:"3" validate:"optional,printascii,min=3,max=100"`
	Icon string `example:"fa-gear" json:"icon" maxLength:"100" minLength:"0" validate:"optional,printascii,min=0,max=100"` // The Icon of the subject. Must be an icon that can be rendered in the browser
}

// Request handler to update a subject
//
// @Summary 	Updates the specified subject [WIP]
// @Description	Updates the subject. Note that all attributes in the body are optional and if omitted, will not be updated
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The id of the subject to update"
// @Param 		payload body UpdateSubjectRequest true "Subject details to update with. All attributes optional!"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "subject not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/update/{id} [patch]
func (m *Controller) UpdateSubject(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to delete a subject
//
// @Summary 	Deletes the given subject [WIP]
// @Description	Deletes the given subject including all subtopics, assets and flashcards.
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The id of the subject to delete"
// @Success		204	{object}		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "subject not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/delete/{id} [delete]
func (m *Controller) DeleteSubject(c *gin.Context) {
	c.JSON(200, nil)
}
