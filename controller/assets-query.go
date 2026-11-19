package controller

import (
	"errors"
	"fmt"

	"github.com/TwiN/logr"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/studyphoenix/backend/httputil"
	"github.com/studyphoenix/backend/model"
)

// Request handler to get a user's flashcards
//
// @Summary 	Get all assets or some
// @Description	Returns a list of user assets, without the actual asset. That needs to be requested separately
// @Tags		assets
// @Produce		json
// @Param 		count query uint false "The maximum number of results to return. Defaults to 25"
// @Param 		offset query uint false "The offset in the results list. Defaults to 0"
// @Success		200	{object}		[]model.Asset
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/assets/search [get]
func (m *Controller) QueryAssets(c *gin.Context) {
	var req model.AssetQuery
	err := c.BindQuery(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse query: ", err))
		c.AbortWithStatus(400)
		return
	}

	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}
	assets, err := m.db.QueryAssets(req, uid)
	if err != nil {
		httputil.NewError(c, 500, errors.New("Failed to process query request"))
		c.Abort()
	} else {
		c.JSON(200, assets)
	}
}
