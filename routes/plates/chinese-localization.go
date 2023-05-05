package plates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
@description 注册汉化区相关路由

	@param router 路由引擎实例
*/
func registerChineseLocalizationPlateRoutes(router *gin.Engine) {
	url := "/plate/chinese-localization"

	router.GET(url, func(context *gin.Context) {
		context.String(http.StatusOK, "plates")
	})
}
