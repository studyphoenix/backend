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

// Request handler to get a subject's tech tree
//
// @Summary 	Get a the tech tree for a given subject
// @Description	Returns the tech tree data structure for rendering the tech tree
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The ID of the subject to retrieve the techtree for"
// @Success		200	{object}		[]model.Subtopic "This is simply a list of subtopics"
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/techtree/get/{id} [get]
func (m *Controller) GetSubjectTechTree(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}

	tree, err := m.db.GetTechTreeForSubject(uint(id), uid)
	if err != nil {
		httputil.NewError(c, 500, errors.New("ERR_TECH_TREE_RETRIEVE"))
		c.Abort()
		return
	}
	c.JSON(200, tree)
}

// Request handler to get a subject's tech tree
//
// @Summary 	Bulk-Creates a tech tree for a subject
// @Description	Creates the tech tree with bulk ingestion
// @Tags		subjects
// @Produce		json
// @Param 		payload body model.CreateSubtopicBulkRequest true "Subtopic details to store"
// @Success		201	{object}		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/techtree/create [put]
func (m *Controller) BulkCreateSubjectTechTree(c *gin.Context) {
	var req model.CreateSubtopicBulkRequest
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

	ok = m.db.BulkCreateTechTreeForSubject(req, uid)
	if ok {
		c.JSON(201, nil)
	} else {
		httputil.NewError(c, 500, errors.New("Cannot create tech tree. See server logs"))
		c.Abort()
		return
	}
}
