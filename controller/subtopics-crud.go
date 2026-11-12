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

// Request handler to create a new subtopic
//
// @Summary 	Create a new subtopic [WIP]
// @Description	Creates a new subtopic linked to the current user
// @Tags		subtopics
// @Produce		json
// @Param 		payload body model.CreateSubtopicRequest true "Subtopic details to store"
// @Success		201	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/create [put]
func (m *Controller) CreateSubtopic(c *gin.Context) {
	var req model.CreateSubtopicRequest
	err := c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create subtopic: ", err))
		c.AbortWithStatus(400)
		return
	}

	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}
	if m.db.CreateSubtopic(req, uid) {
		c.JSON(200, nil)
	} else {
		httputil.NewError(c, 500, errors.New("ERR_CREATE_SUBTOPIC"))
	}
}

// Request handler to get a subtopic
//
// @Summary 	Get the specified subtopic [WIP]
// @Description	Returns the specified subtopic details, including extra metadata
// @Tags		subtopics
// @Produce		json
// @Param 		id path string true "The id of the subtopic to retrieve"
// @Success		200	{object}		model.Subtopic
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/get/{id} [get]
func (m *Controller) GetSubtopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	subtopic, err := m.db.GetSubtopic(uint(id))
	if err != nil {
		c.AbortWithStatus(404)
	}
	c.JSON(200, subtopic)
}

// Request handler to update a subtopic
//
// @Summary 	Updates the specified subtopic [WIP]
// @Description	Updates the subtopic. Note that all attributes in the body are optional and if omitted, will not be updated
// @Tags		subtopics
// @Produce		json
// @Param 		id path string true "The id of the subtopic to update"
// @Param 		payload body model.UpdateSubtopicRequest true "Subtopic details to update with. All attributes optional!"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "subtopic not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/update/{id} [patch]
func (m *Controller) UpdateSubtopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.UpdateSubtopicRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to update subtopic, error: ", err))
		c.Abort()
		return
	}

	if m.db.UpdateSubtopic(uint(id), req) {
		c.JSON(200, nil)
	} else {
		c.JSON(404, nil)
	}
}

// Request handler to delete a subtopic
//
// @Summary 	Deletes the given subtopic [WIP]
// @Description	Deletes the given subtopic including assets and flashcards.
// @Tags		subtopics
// @Produce		json
// @Param 		id path string true "The id of the subtopic to delete"
// @Success		204	{object}		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "subtopic not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/delete/{id} [delete]
func (m *Controller) DeleteSubtopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	if m.db.DeleteSubtopic(uint(id)) {
		c.JSON(200, nil)
	} else {
		c.JSON(404, nil)
	}
}
