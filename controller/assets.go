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

// Request handler to create a new asset
//
// @Summary 	Create a new asset
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
		logr.Debug(fmt.Sprint("Failed to parse body to create asset, error: ", err))
		c.AbortWithStatus(400)
		return
	}

	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}

	id, ok := m.db.CreateAsset(req, uid)
	if ok {
		c.JSON(201, CreatedResponse{
			Id: id,
		})
	} else {
		httputil.NewError(c, 500, errors.New("ERR_CREATE_ASSET"))
		c.Abort()
	}
}

// Request handler to get a asset
//
// @Summary 	Get the specified asset
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
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
// @Summary 	Updates the specified asset
// @Description	Updates the asset. Note that all attributes in the body are optional and if omitted, will not be updated
// @Tags		assets
// @Produce		json
// @Param 		id path string true "The id of the asset to update"
// @Param 		payload body model.UpdateAssetRequest true "Asset details to update with. All attributes optional!"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "asset not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/update/{id} [patch]
func (m *Controller) UpdateAssetDetails(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.UpdateAssetRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to update asset, error: ", err))
		c.Abort()
		return
	}

	if m.db.UpdateAsset(uint(id), req) {
		c.JSON(204, nil)
	} else {
		c.JSON(404, nil)
	}
}

// Request handler to delete a asset
//
// @Summary 	Deletes the given asset
// @Description	Deletes the given asset, including the file stored on disk
// @Tags		assets
// @Produce		json
// @Param 		id path string true "The id of the asset to delete"
// @Success		204	{object}		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "asset not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/delete/{id} [delete]
func (m *Controller) DeleteAsset(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	if m.db.DeleteAsset(uint(id)) {
		c.JSON(204, nil)
	} else {
		c.JSON(404, nil)
	}
}
