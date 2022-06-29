package admin

import (
	v1 "study/app/admin/apis/v1"
	"study/middleware"

	"github.com/gin-gonic/gin"
)

type ApiUserRouter struct{}

func (o *ApiUserRouter) UserRouter(Router *gin.RouterGroup) {
	api := v1.UserController{}
	r := Router.Group("v1").Use(middleware.JwtAuth())
	{
		r.GET("/user/register", api.Register)
	}
}
