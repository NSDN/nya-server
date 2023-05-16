package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取汉化区帖文列表控制器
func GetLocalizationPlate(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"id":        '1',
		"title":     "测试一下过长的帖子标题测试一下过长的帖子标题",
		"thumbnail": "https://static-event.benghuai.com/new_mihoyo_homepage/images/download/cg/origin/2020-09-24.jpg",
	})
}
