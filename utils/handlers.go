package utils

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 打印样品信息
func HandlePrintExample(payload ...any) {
	log.Println(payload...)
}

// 处理请求产生的错误。应当在控制器中调用。
// 封装的是在向客户端返回错误的同时将错误信息打印在服务端的功能。
// 注意此函数不会中断后续请求，如有需要中断请求，请在外部手动执行。
func HandleRequestError(context *gin.Context, httpStatus int, err error) {
	log.Println(err)
	context.String(httpStatus, err.Error())
}

// 处理错误令牌
// 封装的是在向客户端返回错误的同时将错误信息打印在服务端的功能。
// 注意此函数会中断当前请求的后续处理。
func HandleWrongTokenError(context *gin.Context) {
	HandleRequestError(
		context,
		http.StatusUnauthorized,
		errors.New(Messages.AUTHORIZE_FAILED_WRONG_TOKEN),
	)

	context.Abort()
}
