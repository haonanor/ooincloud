package initialize

import (
	"io/fs"
	"net/http"
	"study/middleware"
	"study/recourse"
	"study/router"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Static())
	stripped, err := fs.Sub(recourse.StaticPath, "dist/assets")
	if err != nil {
		panic(err)
	}
	Router.StaticFS("/assets", http.FS(stripped))
	adminRouter := router.RouterGroupApp.Admin
	Group := Router.Group("api")
	{
		adminRouter.UserRouter(Group)
		adminRouter.FileRouter(Group)
	}

	return Router
}
