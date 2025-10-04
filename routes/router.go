package routes

import (
	"net/http"

	"github.com/NSDN/nya-server/context"
	"github.com/NSDN/nya-server/middleware"
	"github.com/gin-gonic/gin"
)

// 初始化路由
func SetupRouter(context *context.AppContext) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORS())
	router.Use(middleware.ValidateAuthorization)

	// HTTP 连接测试 API
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	context.APIRouter = router.Group("/api")

	registerPlatesRoutes(context)
	registerAuthorizationRoutes(context)
	registerTopicRoutes(context)

	return router
}
