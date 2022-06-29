package initialize

import (
	"study/router"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	adminRouter := router.RouterGroupApp.Admin
	Group := Router.Group("api")
	{
		adminRouter.UserRouter(Group)
	}
	return Router
}
