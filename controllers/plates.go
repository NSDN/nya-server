package controllers

import (
	"net/http"

	"github.com/NSDN/nya-server/services"
	"github.com/NSDN/nya-server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建版块列表 - 控制器
func InitPlateList(db *gorm.DB) {
	services.InitPlateList(db)
}

// 获取版块列表 - 控制器
func GetPlateList(context *gin.Context, db *gorm.DB) {
	list, err := services.GetPlateList(db)

	if err != nil {
		utils.HandleRequestError(context, http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, list)
}
