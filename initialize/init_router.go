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
	// Router.AppEngine.Header
	Router.Use(middleware.Cors())
	// Router.Header("Access-Control-Allow-Origin", "*")
	// Router.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	// Router.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	// o.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	// o.Header("Access-Control-Allow-Credentials", "true")
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
