package controller

import "github.com/gin-gonic/gin"

// Request handler to upload a new asset
//
// @Summary 	Upload file to a given asset [WIP]
// @Description	Upload a file to asset with id `{id}`, overriding any previously existing files irreversibly
// @Tags		assets
// @Produce		json
// @Accept		multipart/form-data
// @Param 		id path string true "The id of the asset to upload it to"
// @Param 		payload formData file true "The file to upload here"
// @Success		201	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/upload/{id} [put]
func (m *Controller) UploadAsset(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to download an asset
//
// @Summary 	Download the file of a given asset [WIP]
// @Description	Download the file of asset with id `{id}`
// @Tags		assets
// @Param 		id path string true "The id of the asset to download"
// @Success		200	{file}			nil "The requested file is in the body"
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/download/{id} [get]
func (m *Controller) DownloadAsset(c *gin.Context) {
	c.JSON(200, nil)
}
