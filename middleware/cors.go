package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,Upload-Length,Upload-Offset,Upload-Name,Upload-Length")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,PATCH,DELETE,HEAD")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type,Upload-Length,Upload-Offset,Upload-Name,Upload-Length")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}
