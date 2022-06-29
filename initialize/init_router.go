package initialize

import (
	"net/http"
	"study/assets"
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
	Router.StaticFS("/ui", http.FS(assets.Static))
	return Router
}
