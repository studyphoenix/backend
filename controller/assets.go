package controller

import (
	"github.com/gin-gonic/gin"
)

type CreateAssetRequest struct {
	Name string `example:"Analysis" json:"name" maxLength:"250" minLength:"3"`
	Icon string `example:"fa-gear" json:"icon" maxLength:"250" minLength:"0"` // The Icon of the asset. Must be an icon that can be rendered in the browser
}

// Request handler to create a new asset
//
// @Summary 	Create a new asset [WIP]
// @Description	Creates a new asset linked to the current user
// @Tags		assets
// @Produce		json
// @Param 		payload body CreateAssetRequest true "Asset details to store"
// @Success		201	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/create [put]
func (m *Controller) CreateAsset(c *gin.Context) {
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
	c.JSON(200, nil)
}

type UpdateAssetRequest struct {
	Name *string `example:"Analysis" json:"name" maxLength:"250" minLength:"3"`
	Icon *string `example:"fa-gear" json:"icon" maxLength:"250" minLength:"0"` // The Icon of the asset. Must be an icon that can be rendered in the browser
}

// Request handler to update a asset
//
// @Summary 	Updates the specified asset [WIP]
// @Description	Updates the asset. Note that all attributes in the body are optional and if omitted, will not be updated
// @Tags		assets
// @Produce		json
// @Param 		id path string true "The id of the asset to update"
// @Param 		payload body UpdateAssetRequest true "Asset details to update with. All attributes optional!"
// @Success		200	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure 	404 {object} 		nil "asset not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/update/{id} [patch]
func (m *Controller) UpdateAssetDeteails(c *gin.Context) {
	c.JSON(200, nil)
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
	c.JSON(200, nil)
}
