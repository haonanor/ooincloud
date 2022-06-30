package router

import (
	"study/router/admin"
)

// var AppRouters = make([]func(*gin.RouterGroup), 0)

// func AddRouters(actionName func(*gin.RouterGroup)) {
// 	AppRouters = append(AppRouters, actionName)
// }

// func InitRouter() *gin.Engine {
// 	var Router = gin.Default()
// 	Group := Router.Group("api")
// 	{
// 		for _, f := range AppRouters {
// 			f(Group)
// 		}
// 	}
// 	fmt.Println(AppRouters)
// 	return Router
// }

type RouterGroup struct {
	Admin admin.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
