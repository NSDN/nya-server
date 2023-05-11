package plates

import (
	"net/http"

	services "github.com/NSDN/nya-server/services/plates"
	"github.com/gin-gonic/gin"
)

// 创建版块列表 - 控制器
func InitPlateList() {
	services.InitPlateList()
}

// 获取版块列表 - 控制器
func GetPlateList(context *gin.Context) {
	list, err := services.GetPlateList()

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	}

	context.JSON(http.StatusOK, list)
}
