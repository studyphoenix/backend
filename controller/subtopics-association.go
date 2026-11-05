package controller

import "github.com/gin-gonic/gin"

// Request handler to (de)associate subtopic(s)
//
// @Summary 	(de)associate subtopic(s) [WIP]
// @Description	(de)associates the subtopic(s) depending on the value of "add" to this subtopic
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The id of the subtopic to update"
// @Param 		payload body AssociationRequest true "The associations to add or remove"
// @Success		200	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "subject or id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/subtopics/{id} [patch]
func (m *Controller) AssociateSubtopicToSubtopic(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to (de)associate asset(s)
//
// @Summary 	(de)associate asset(s) [WIP]
// @Description	(de)associates the asset(s) depending on the value of "add" to this subtopic
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The id of the subtopic to update"
// @Param 		payload body AssociationRequest true "The associations to add or remove"
// @Success		200	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	400 {object} 		nil "subject or id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subtopics/assets/{id} [patch]
func (m *Controller) AssociateAssetToSubtopic(c *gin.Context) {
	c.JSON(200, nil)
}
