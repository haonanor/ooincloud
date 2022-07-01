package middleware

import (
	"strings"
	"study/recourse"

	"github.com/gin-gonic/gin"
)

func Static() gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path

		if strings.HasPrefix(path, "/api") || strings.HasPrefix(path, "/assets") {
			c.Next()
			return
		}
		c.Header("Content-Type", "text/html")
		c.String(200, recourse.MainIndexPath)
		c.Abort()

	}
}
