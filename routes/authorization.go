package routes

import (
	"github.com/NSDN/nya-server/controllers"
	"github.com/gin-gonic/gin"
)

// 注册认证相关路由
func registerAuthorizationRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)
}
