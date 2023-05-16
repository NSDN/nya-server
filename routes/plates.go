package routes

import (
	"github.com/NSDN/nya-server/controllers"
	"github.com/gin-gonic/gin"
)

// 注册版块相关路由。
// 应当在初始化路由的函数中调用它，并传入一个路由引擎实例的指针。
func registerPlatesRoutes(router *gin.Engine) {
	// 创建模版列表
	controllers.InitPlateList()

	router.GET("/plates", controllers.GetPlateList)
	registerLocalizationPlateRoutes(router)
}
