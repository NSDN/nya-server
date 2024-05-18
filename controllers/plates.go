package controllers

import (
	"net/http"

	"github.com/NSDN/nya-server/services"
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

// TODO: 等创建文章功能完成后应删除此函数
//
// 创建假帖文列表 - 控制器
func InitDummyArticles() {
	services.InitDummyArticles()
}

// 获取帖文列表 - 控制器
//
// 从路由上下文中找出路由参数（版块路由名称）传递给服务，服务会根据路由参数返回不同版块的文章。
func GetArticles(context *gin.Context) {
	plate := context.Params.ByName("plate")

	list, err := services.GetArticles(plate)

	if err != nil {
		utils.HandleRequestError(context, http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, list)
}
