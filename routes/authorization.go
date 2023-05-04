package routes

import (
	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/controllers"
	"github.com/gin-gonic/gin"
)

// 注册认证相关路由
func registerAuthorizationRoutes(router *gin.Engine) {
	router.POST(configs.API_REGISTER, controllers.Register)
	router.POST(configs.API_LOGIN, controllers.Login)
	router.GET(configs.API_GET_USER_INFO, controllers.GetUserInfo)
}
