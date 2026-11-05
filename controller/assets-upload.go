package controller

import "github.com/gin-gonic/gin"

// Request handler to create a new asset
//
// @Summary 	Upload file to a given asset [WIP]
// @Description	Upload a file to asset with id `{id}`, overriding any previously existing files irreversibly
// @Tags		assets
// @Produce		json
// @Accept		multipart/form-data
// @Param 		payload formData file true "The file to upload here"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/upload/{id} [put]
func (m *Controller) UploadAsset(c *gin.Context) {
	c.JSON(200, nil)
}
