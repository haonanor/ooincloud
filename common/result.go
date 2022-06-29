package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Default = &Response{}

func Ok(c *gin.Context, data interface{}, message string) {
	res := Default.Clone()

	res.SetCode(1)
	res.SetData(data)
	res.SetMessage(message)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func Error(c *gin.Context, code int32, message string) {
	res := Default.Clone()

	res.SetCode(code)
	res.SetMessage(message)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
