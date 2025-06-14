package controllers

import (
	"errors"
	"net/http"

	"github.com/NSDN/nya-server/context"
	"github.com/NSDN/nya-server/models"
	"github.com/NSDN/nya-server/services"
	"github.com/NSDN/nya-server/utils"
	"github.com/gin-gonic/gin"
)

// 获取帖子列表 - 控制器
func GetTopics(appContext *context.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		service := services.NewTopicService(appContext)
		plateID := context.Query("plateId")

		list, err := service.GetTopics(plateID)

		if err != nil {
			utils.HandleRequestError(context, http.StatusInternalServerError, err)
			return
		}

		context.JSON(http.StatusOK, list)
	}
}

// 获取帖文列表 - 控制器
//
// 从路由上下文中找出路由参数（版块路由名称）传递给服务，服务会根据路由参数返回不同版块的文章。
func GetArticles(context *gin.Context) {
	plate := context.Params.ByName("plate")

	list, err := services.GetTopicsOld(plate)

	if err != nil {
		utils.HandleRequestError(context, http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, list)
}

// 获取帖子楼层列表 - 控制器
//
// 从路由上下文中找出路由参数（版块路由名称）传递给服务，
// 服务会根据路由参数返回不同文章的楼层列表。
func GetTopicFloors(context *gin.Context) {
	topic := context.Params.ByName("topic")

	list, err := services.GetTopicFloors(topic)

	if err != nil {
		utils.HandleRequestError(context, http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, list)
}

// 创建文章 - 控制器
func CreateArticle(context *gin.Context) {
	var request models.NewTopicRequestData
	err := context.ShouldBindJSON(&request)

	if err != nil {
		utils.HandleRequestError(
			context,
			http.StatusBadRequest,
			errors.New(utils.Messages.ARTICLE_FAILED_BAD_CONTECT),
		)
		return
	}

	succeed, err := services.CreateTopic(&request)

	if err != nil {
		utils.HandleRequestError(context, http.StatusForbidden, err)
		return
	}

	context.JSON(http.StatusOK, succeed)
}
