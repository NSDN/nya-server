package routes

import (
	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/context"
	"github.com/NSDN/nya-server/controllers"
)

// 注册认证相关路由
func registerAuthorizationRoutes(context *context.AppContext) {
	controller := controllers.NewAuthorizationController(context)

	registerUserRegisterRoute(context, controller)
	registerLoginRoute(context, controller)
	registerUserInfoRoute(context, controller)
}

func registerUserRegisterRoute(
	context *context.AppContext,
	controller *controllers.AuthorizationController,
) {
	context.APIRouter.POST(configs.API_REGISTER, controller.Register)
}

func registerLoginRoute(
	context *context.AppContext,
	controller *controllers.AuthorizationController,
) {
	context.APIRouter.POST(configs.API_LOGIN, controller.Login)
}

func registerUserInfoRoute(
	context *context.AppContext,
	controller *controllers.AuthorizationController,
) {
	context.APIRouter.GET(configs.API_GET_USER_INFO, controller.GetUserInfo)
}
