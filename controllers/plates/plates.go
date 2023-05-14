package plates

import (
	"net/http"

	services "github.com/NSDN/nya-server/services/plates"
	"github.com/NSDN/nya-server/utils"
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
		utils.HandleRequestError(context, http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, list)
}
