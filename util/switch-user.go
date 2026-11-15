package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SwitchUser(r *gin.Engine, prod bool) {
	// Can set UID via /testing/uid?userid=<new id>
	r.GET("/testing/uid", func(c *gin.Context) {
		sess := sessions.Default(c)
		sess.Set("userid", c.Query("userid"))
	})
}
