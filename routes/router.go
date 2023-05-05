package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var _router *gin.Engine

func SetupRouter() *gin.Engine {
	_router = gin.Default()

	pingTest := func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	}

	_router.GET("/ping", pingTest)

	return _router
}
