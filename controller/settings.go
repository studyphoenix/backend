package controller

import "github.com/gin-gonic/gin"

// Request handler to get a user's config
//
// @Summary 	Get the user's config [WIP]
// @Description	Returns the user's stored configuration
// @Tags		user
// @Produce		json
// @Success		200	{object}		model.Settings
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/user/config [get]
func (m *Controller) GetConfig(c *gin.Context) {
	c.JSON(200, nil)
}

// Request handler to update a user's config
//
// @Summary 	Update the user's config [WIP]
// @Description	Update the user's config using the provided details. All of them are optional
// @Tags		user
// @Produce		json
// @Param 		payload body model.UpdateUserConfigRequest true "Configuration Details to update"
// @Success		200	{object}		model.Settings
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/user/config [patch]
func (m *Controller) StoreConfig(c *gin.Context) {
	c.JSON(200, nil)
}
