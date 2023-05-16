package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

// 处理请求产生的错误。应当在控制器中调用。
// 封装的是在向客户端返回错误的同时将错误信息打印在服务端的功能。
// 注意此函数不会中断后续请求，如有需要，请在外部手动执行。
func HandleRequestError(context *gin.Context, httpStatus int, err error) {
	log.Println(err)
	context.String(httpStatus, err.Error())
}
