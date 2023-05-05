package main

import (
	"log"
	"os"

	"github.com/NSDN/nya-server/repositories"
	"github.com/NSDN/nya-server/router"
	"github.com/NSDN/nya-server/utils"
	"github.com/joho/godotenv"
)

func init() {
	// 从 .env 文件中加载环境变量
	err := godotenv.Load()

	if err != nil {
		log.Fatal(utils.Messages.ENVIRONMENT_ERROR_NOT_FOUND)
	}
}

func main() {
	r := router.SetupRouter()
	repositories.SetupDatabase()

	APPLICATION_PORT := "APPLICATION_PORT"
	port := os.Getenv(APPLICATION_PORT)

	if port == "" {
		log.Fatal(utils.Messages.ENVIRONMENT_ERROR_NEED_ONE(APPLICATION_PORT))
	}

	r.Run(port)
}
