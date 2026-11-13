package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hci-group-15/backend/httputil"
)

// Request handler to get a subject's tech tree
//
// @Summary 	Get a the tech tree for a given subject [WIP]
// @Description	Returns the tech tree data structure for rendering the tech tree
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The ID of the subject to retrieve the techtree for"
// @Success		200	{object}		[]model.Subtopic "This is simply a list of subtopics"
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/techtree/{id} [get]
func (m *Controller) GetSubjectTechTree(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	tree, err := m.db.GetTechTreeForSubject(uint(id))
	if err != nil {
		httputil.NewError(c, 500, errors.New("ERR_TECH_TREE_RETRIEVE"))
		c.Abort()
		return
	}
	c.JSON(200, tree)
}
