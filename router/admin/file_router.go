package admin

import (
	v1 "study/app/admin/apis/v1"

	"github.com/gin-gonic/gin"
)

type ApiFileRouter struct{}

func (o *ApiFileRouter) FileRouter(Router *gin.RouterGroup) {
	api := v1.FileController{}
	r := Router.Group("v1")
	{
		r.POST("/file/upload", api.Post)
		r.PATCH("/file/upload", api.Patch)
	}
}
