package controllers

import (
	"net/http"

	"github.com/NSDN/nya-server/context"
	"github.com/NSDN/nya-server/services"
	"github.com/NSDN/nya-server/utils"
	"github.com/gin-gonic/gin"
)

// 创建版块列表 - 控制器
func InitPlateList(context *context.AppContext) {
	plateService := services.NewPlateService(context)
	plateService.InitPlateList()
}

// 获取版块列表 - 控制器
func GetPlates(appContext *context.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		service := services.NewPlateService(appContext)
		list, err := service.GetPlates()

		if err != nil {
			utils.HandleRequestError(context, http.StatusInternalServerError, err)
			return
		}

		context.JSON(http.StatusOK, list)
	}
}
