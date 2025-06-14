package routes

import (
	"github.com/NSDN/nya-server/context"
	"github.com/NSDN/nya-server/controllers"
)

// 注册版块相关路由。
// 应当在初始化路由的函数中调用它，并传入一个路由引擎实例的指针。
func registerPlatesRoutes(context *context.AppContext) {
	registerQueryPlateListRoute(context)
}

/** 注册请求版块列表的路由 */
func registerQueryPlateListRoute(appContext *context.AppContext) {
	// 因为版块列表是保存在后端的，所以需要由后端初始化
	controllers.InitPlateList(appContext)

	appContext.APIRouter.GET("/plates", controllers.GetPlates(appContext))
}
