package util

import (
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SwitchUser(r *gin.Engine, prod bool) {
	// Can set UID via /testing/uid?userid=<new id>
	r.GET("/testing/uid", func(c *gin.Context) {
		q, err := strconv.ParseUint(c.Query("userid"), 10, 32)
		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		sess := sessions.Default(c)
		sess.Set("userid", uint(q))
		c.JSON(200, gin.H{
			"userid": sess.Get("userid"),
		})
		sess.Save()
	})
}
