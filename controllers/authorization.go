package controllers

import (
	"errors"
	"net/http"

	"github.com/NSDN/nya-server/models"
	"github.com/NSDN/nya-server/services"
	"github.com/NSDN/nya-server/utils"
	"github.com/gin-gonic/gin"
)

// 登入 - 控制器
func Login(context *gin.Context) {
	var info models.LoginInfo

	// 绑定登入信息
	if err := context.ShouldBindJSON(&info); err != nil {
		utils.HandleRequestError(
			context,
			http.StatusBadRequest,
			errors.New(utils.Messages.AUTHORIZE_FAILED_BAD_PARAMETER),
		)

		return
	}

	// 验证参数
	if info.Username == "" {
		utils.HandleRequestError(
			context,
			http.StatusBadRequest,
			errors.New(utils.Messages.AUTHORIZE_FAILED_MISSING_PARAMETER("用户名")),
		)

		return
	}

	if info.Password == "" {
		utils.HandleRequestError(
			context,
			http.StatusBadRequest,
			errors.New(utils.Messages.AUTHORIZE_FAILED_MISSING_PARAMETER("密码")),
		)

		return
	}

	// 调用登入服务
	if err := services.Login(info); err != nil {
		utils.HandleRequestError(context, http.StatusForbidden, err)
		return
	}

	context.String(http.StatusOK, "called")
}
