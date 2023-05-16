package routes

import (
	"net/http"

	"github.com/NSDN/nya-server/middleware"
	"github.com/gin-gonic/gin"
)

// 初始化路由
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORS())

	// HTTP 连接测试 API
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	registerPlatesRoutes(router)
	registerAuthorizationRoutes(router)

	return router
}
