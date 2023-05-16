package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册汉化区相关路由。
// 应当在注册版块路由的函数中调用它并传入一个路由引擎实例。
func registerLocalizationPlateRoutes(router *gin.Engine) {
	url := "/plate/localization"

	router.GET(url, func(context *gin.Context) {
		context.String(http.StatusOK, "plates")
	})
}
