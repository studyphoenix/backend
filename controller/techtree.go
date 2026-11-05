package controller

import "github.com/gin-gonic/gin"

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
	c.JSON(200, nil)
}
