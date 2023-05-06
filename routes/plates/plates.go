package plates

import (
	plateControllers "github.com/NSDN/nya-server/controllers/plates"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	router.GET("/plates", plateControllers.GetPlateList)
	registerLocalizationPlateRoutes(router)
}
