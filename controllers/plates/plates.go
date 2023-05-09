package plates

import (
	"net/http"

	platesServices "github.com/NSDN/nya-server/services/plates"
	"github.com/gin-gonic/gin"
)

func GetPlateList(context *gin.Context) {
	list := platesServices.GetPlateList()
	context.JSON(http.StatusOK, list)
}
