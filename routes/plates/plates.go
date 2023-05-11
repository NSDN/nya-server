package plates

import (
	controllers "github.com/NSDN/nya-server/controllers/plates"
	"github.com/gin-gonic/gin"
)

// 注册版块相关路由
func Register(router *gin.Engine) {
	// 创建模版列表
	controllers.InitPlateList()

	router.GET("/plates", controllers.GetPlateList)
	registerLocalizationPlateRoutes(router)
}
