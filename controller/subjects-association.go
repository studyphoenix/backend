package controller

import (
	"fmt"
	"strconv"

	"github.com/TwiN/logr"
	"github.com/gin-gonic/gin"
	"github.com/studyphoenix/backend/model"
)

// Request handler to associate asset(s)
//
// @Summary 	associate asset(s)
// @Description	associates the specified asset(s) to this subject
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The id of the subject to update"
// @Param 		payload body model.AssociationRequest true "The associations to add"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "subject or id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/associate/assets/{id} [patch]
func (m *Controller) AssociateAssetToSubject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.AssociationRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create association: ", err))
		c.AbortWithStatus(400)
		return
	}

	for _, assetId := range req {
		m.db.AssociateAssetToSubject(uint(id), assetId)
	}

	c.JSON(204, nil)
}

// Request handler to associate asset(s)
//
// @Summary 	deassociate asset(s)
// @Description	deassociates the specified asset(s) to this subject
// @Tags		subjects
// @Produce		json
// @Param 		id path string true "The id of the subject to update"
// @Param 		payload body model.AssociationRequest true "The associations to remove"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "subject or id not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/subjects/deassociate/assets/{id} [patch]
func (m *Controller) DeassociateAssetFromSubject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.AssociationRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create association: ", err))
		c.AbortWithStatus(400)
		return
	}

	for _, assetId := range req {
		m.db.DeassociateAssetFromSubject(uint(id), assetId)
	}

	c.JSON(204, nil)
}
