package routes

import (
	"github.com/NSDN/nya-server/controllers"
	"github.com/gin-gonic/gin"
)

// 注册版块相关路由。
// 应当在初始化路由的函数中调用它，并传入一个路由引擎实例的指针。
func registerPlatesRoutes(router *gin.Engine) {
	registerQueryPlateListRoute(router)
}

/** 注册请求版块列表的路由 */
func registerQueryPlateListRoute(router *gin.Engine) {
	// 因为版块列表是保存在后端的，所以需要由后端初始化
	controllers.InitPlateList()
	router.GET("/plates", controllers.GetPlateList)
}
