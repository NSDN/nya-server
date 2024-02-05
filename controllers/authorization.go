package controllers

import (
	"errors"
	"net/http"

	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/models"
	"github.com/NSDN/nya-server/services"
	"github.com/NSDN/nya-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 注册 - 控制器
func Register(context *gin.Context) {
	var info models.RegisterInfo

	// 从请求体中获取注册信息
	if err := context.ShouldBindJSON(&info); err != nil {
		utils.HandleRequestError(
			context,
			http.StatusBadRequest,
			errors.New(utils.Messages.AUTHORIZE_FAILED_BAD_PARAMETER),
		)
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

	// 调用注册服务
	token, err := services.Login(
		models.LoginInfo{
			Username: info.Username,
			Password: info.Password,
		},
	)

	if err != nil {
		utils.HandleRequestError(context, http.StatusForbidden, err)
		return
	}

	context.JSON(http.StatusOK, models.Token{
		AccessToken: token,
	})
}

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
	token, err := services.Login(info)

	if err != nil {
		utils.HandleRequestError(context, http.StatusForbidden, err)
		return
	}

	context.JSON(http.StatusOK, models.Token{
		AccessToken: token,
	})
}

// 获取用户信息 - 控制器
func GetUserInfo(context *gin.Context) {
	claims, exist := context.Get(configs.CONTEXT_KEY_CLAIMS)

	if !exist {
		utils.HandleWrongTokenError(context)
		return
	}

	uid := claims.(jwt.MapClaims)["uid"].(string)

	// 调用用户信息
	user, err := services.GetUserInfo(uid)

	if err != nil {
		utils.HandleRequestError(context, http.StatusForbidden, err)
		return
	}

	context.JSON(http.StatusOK, user)
}
