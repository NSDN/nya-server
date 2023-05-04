package middleware

import (
	"github.com/NSDN/nya-server/configs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, configs.HTTP_HEADER_AUTHORIZATION)
	return cors.New(config)
}
