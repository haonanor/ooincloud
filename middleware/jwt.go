package middleware

import (
	"study/common"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			common.Error(c, 400001, "非法登录")
		}
	}
}
