package controller

import "github.com/gin-gonic/gin"

// Request handler to associate subtopic(s)
//
// @Summary 	associate subtopic(s) [WIP]
// @Description	associates the subtopic(s) from this subject
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The id of the subject to update"
// @Param 		payload body model.AssociationRequest true "The associations to add"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "subject or id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/associate/subtopics/{id} [patch]
func (m *Controller) AssociateSubtopicToSubject(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to deassociate subtopic(s)
//
// @Summary 	deassociate subtopic(s) [WIP]
// @Description	deassociates the specified subtopic(s) from this subject
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The id of the subject to update"
// @Param 		payload body model.AssociationRequest true "The associations to remove"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "subject or id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/deassociate/subtopics/{id} [patch]
func (m *Controller) DeassociateSubtopicToSubject(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to associate asset(s)
//
// @Summary 	associate asset(s) [WIP]
// @Description	associates the specified asset(s) to this subject
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The id of the subject to update"
// @Param 		payload body model.AssociationRequest true "The associations to add"
// @Success		200	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	400 {object} 		nil "subject or id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/associate/assets/{id} [patch]
func (m *Controller) AssociateAssetToSubject(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to associate asset(s)
//
// @Summary 	deassociate asset(s) [WIP]
// @Description	deassociates the specified asset(s) to this subject
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The id of the subject to update"
// @Param 		payload body model.AssociationRequest true "The associations to remove"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	400 {object} 		nil "subject or id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/deassociate/assets/{id} [patch]
func (m *Controller) DeassociateAssetToSubject(c *gin.Context) {
	c.JSON(200, nil)
}
