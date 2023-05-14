package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

// 处理请求产生的错误
func HandleRequestError(context *gin.Context, status int, err error) {
	log.Println(err)
	context.String(status, err.Error())
}
