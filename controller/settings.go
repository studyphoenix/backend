package controller

import (
	"errors"
	"fmt"

	"github.com/TwiN/logr"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hci-group-15/backend/httputil"
	"github.com/hci-group-15/backend/model"
)

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
	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}
	settings, err := m.db.GetSettings(uint(uid))
	if err != nil {
		httputil.NewError(c, 500, errors.New("ERR_SETTINGS_GET"))
		return
	}
	c.JSON(200, settings)
}

// Request handler to update a user's config
//
// @Summary 	Update the user's config [WIP]
// @Description	Update the user's config using the provided details. All of them are optional
// @Tags		user
// @Produce		json
// @Param 		payload body model.UpdateUserConfigRequest true "Configuration Details to update"
// @Success		200	{object}		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/user/config [patch]
func (m *Controller) StoreConfig(c *gin.Context) {
	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}
	var req model.UpdateUserConfigRequest
	err := c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to update subject, error: ", err))
		c.Abort()
		return
	}

	if m.db.StoreSettings(uint(uid), req) {
		c.JSON(200, nil)
	} else {
		c.JSON(404, nil)
	}
}
