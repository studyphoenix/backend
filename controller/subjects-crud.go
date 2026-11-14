package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/TwiN/logr"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/studyphoenix/backend/httputil"
	"github.com/studyphoenix/backend/model"
)

// Request handler to create a new subject
//
// @Summary 	Create a new subject
// @Description	Creates a new subject linked to the current user and returns the id of it
// @Tags		subjects
// @Produce		json
// @Param 		payload body model.CreateSubjectRequest true "Subject details to store"
// @Success		201	{object}		CreatedResponse
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/create [put]
func (m *Controller) CreateSubject(c *gin.Context) {
	var req model.CreateSubjectRequest
	err := c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create subject: ", err))
		c.AbortWithStatus(400)
		return
	}

	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}
	id, ok := m.db.CreateSubject(req, uid)
	if ok {
		c.JSON(201, CreatedResponse{
			Id: id,
		})
	} else {
		httputil.NewError(c, 500, errors.New("ERR_CREATE_SUBJECT"))
		c.Abort()
	}
}

// Request handler to get a subject
//
// @Summary 	Get the specified subject
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	subject, err := m.db.GetSubject(uint(id))
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, subject)
}

// Request handler to update a subject
//
// @Summary 	Updates the specified subject
// @Description	Updates the subject. Note that all attributes in the body are optional and if omitted, will not be updated
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The id of the subject to update"
// @Param 		payload body model.UpdateSubjectRequest true "Subject details to update with. All attributes optional!"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "subject not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/update/{id} [patch]
func (m *Controller) UpdateSubject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.UpdateSubjectRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to update subject, error: ", err))
		c.Abort()
		return
	}

	if m.db.UpdateSubject(uint(id), req) {
		c.JSON(204, nil)
	} else {
		c.JSON(404, nil)
	}
}

// Request handler to delete a subject
//
// @Summary 	Deletes the given subject
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	if m.db.DeleteSubject(uint(id)) {
		c.JSON(204, nil)
	} else {
		c.JSON(404, nil)
	}
}
