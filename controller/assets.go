package controller

import (
	"fmt"
	"strconv"

	"github.com/TwiN/logr"
	"github.com/gin-gonic/gin"
	"github.com/hci-group-15/backend/model"
)

// Request handler to create a new asset
//
// @Summary 	Create a new asset [WIP]
// @Description	Creates a new asset linked to the current user
// @Tags		assets
// @Produce		json
// @Param 		payload body model.CreateAssetRequest true "Asset details to store"
// @Success		201	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/create [put]
func (m *Controller) CreateAsset(c *gin.Context) {
	var req model.CreateAssetRequest
	err := c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprintf("Failed to parse body for update product, error: ", err))
		c.Abort()
		return
	}
	m.db.CreateAsset(req)

	c.JSON(200, nil)
}

// Request handler to get a asset
//
// @Summary 	Get the specified asset [WIP]
// @Description	Returns the specified asset details, including extra metadata
// @Tags		assets
// @Produce		json
// @Param 		id path string true "The id of the asset to retrieve"
// @Success		200	{object}		model.Asset
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/get/{id} [get]
func (m *Controller) GetAssetDetails(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	asset, err := m.db.GetAsset(uint(id))
	if err != nil {
		c.AbortWithStatus(404)
	}
	c.JSON(200, asset)
}

// Request handler to update a asset
//
// @Summary 	Updates the specified asset [WIP]
// @Description	Updates the asset. Note that all attributes in the body are optional and if omitted, will not be updated
// @Tags		assets
// @Produce		json
// @Param 		id path string true "The id of the asset to update"
// @Param 		payload body model.UpdateAssetRequest true "Asset details to update with. All attributes optional!"
// @Success		200	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "asset not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/update/{id} [patch]
func (m *Controller) UpdateAssetDetails(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.UpdateAssetRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprintf("Failed to parse body for update product, error: ", err))
		c.Abort()
		return
	}

	if m.db.UpdateAsset(uint(id), req) {
		c.JSON(200, nil)
	} else {
		c.JSON(404, nil)
	}
}

// Request handler to delete a asset
//
// @Summary 	Deletes the given asset [WIP]
// @Description	Deletes the given asset, including the file stored on disk
// @Tags		assets
// @Produce		json
// @Param 		id path string true "The id of the asset to delete"
// @Success		200	{object}		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "asset not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/delete/{id} [delete]
func (m *Controller) DeleteAsset(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	if m.db.DeleteAsset(uint(id)) {
		c.JSON(200, nil)
	} else {
		c.JSON(404, nil)
	}
}
