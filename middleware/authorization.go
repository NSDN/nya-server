package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 查验认证请求头
func ValidateAuthorization(context *gin.Context) {
	authorization := context.GetHeader(configs.HTTP_HEADER_AUTHORIZATION)

	// 如果没有请求头中无认证信息就以游客身份继续
	if authorization == "" {
		context.Next()
		return
	}

	// 如果有认证信息则验证认证信息是否正确
	tokenString := getTokenHeader(authorization)

	if tokenString == "" {
		hanldeWrongToken(context)
		return
	}

	token, err := validateToken(tokenString)

	// 如果令牌验证成功，则向上下文中填入声明集以便于后续使用
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		context.Set(configs.CONTEXT_KEY_CLAIMS, claims)
		context.Next()
	} else {
		log.Println(err)
		hanldeWrongToken(context)
	}
}

// 从请求头中获取令牌字符串
func getTokenHeader(authorization string) (token string) {
	// 检查是否为 Bearer 认证
	isBearer := strings.HasPrefix(authorization, configs.AUTHENTICATION_TYPE)

	if isBearer {
		return strings.TrimPrefix(authorization, configs.AUTHENTICATION_TYPE)
	}

	return ""
}

// 查验令牌
func validateToken(tokenString string) (token *jwt.Token, err error) {
	// 从字符串中解析令牌
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 判断令牌的方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		var tokenKey = utils.GetENV(configs.ENV_TOKEN_KEY)
		return []byte(tokenKey), nil
	})
}

// 处理错误令牌
func hanldeWrongToken(context *gin.Context) {
	utils.HandleRequestError(
		context,
		http.StatusUnauthorized,
		errors.New(utils.Messages.AUTHORIZE_FAILED_WRONG_TOKEN),
	)

	context.Abort()
}
