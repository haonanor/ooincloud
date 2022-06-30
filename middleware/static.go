package middleware

import (
	"net/http"
	"strings"
	"study/recourse"

	"github.com/gin-gonic/gin"
)

func Static() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

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
