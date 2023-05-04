package routes

import (
	"net/http"

	"github.com/NSDN/nya-server/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 初始化路由
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORS())
	router.Use(middleware.ValidateAuthorization)

	// HTTP 连接测试 API
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	apiRouter := router.Group("/api")

	registerPlatesRoutes(apiRouter, db)
	// registerAuthorizationRoutes(router)
	// registerTopicRoutes(router)

	return router
}
