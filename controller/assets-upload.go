package controller

import (
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hci-group-15/backend/model"
)

// Request handler to upload a new asset
//
// @Summary 	Upload file to a given asset
// @Description	Upload a file to asset with id `{id}`, overriding any previously existing files irreversibly
// @Tags		assets
// @Produce		json
// @Accept		multipart/form-data
// @Param 		id path string true "The id of the asset to upload it to"
// @Param 		payload formData file true "The file to upload here"
// @Success		201	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		409	{object}		nil "Expected file type didn't match actual file type"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/upload/{id} [put]
func (m *Controller) UploadAsset(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	asset, err := m.db.GetAsset(uint(id))
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	if model.MEDIA_TYPE_MAP[asset.MediaType] != filepath.Ext(file.Filename) {
		c.AbortWithStatus(409)
		return
	}

	dest := filepath.Join(m.db.Config.Datadir, strconv.FormatUint(id, 10), filepath.Ext(file.Filename))
	c.SaveUploadedFile(file, dest)

	c.JSON(201, nil)
}

// Request handler to download an asset
//
// @Summary 	Download the file of a given asset
// @Description	Download the file of asset with id `{id}`
// @Tags		assets
// @Param 		id path string true "The id of the asset to download"
// @Success		200	{file}			nil "The requested file is in the body"
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/download/{id} [get]
func (m *Controller) DownloadAsset(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	asset, err := m.db.GetAsset(uint(id))
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	src := filepath.Join(m.db.Config.Datadir, strconv.FormatUint(id, 10), "."+model.MEDIA_TYPE_MAP[asset.MediaType])
	c.File(src)
}
