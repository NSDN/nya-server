package utils

import (
	"log"
	"os"
)

// 从 .env 文件中获取环境变量
func GetENV(envName string) string {
	value := os.Getenv(envName)

	if value == "" {
		log.Fatal(Messages.ENVIRONMENT_ERROR_NEED_ONE(envName))
	}

	return value
}
