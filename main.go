package main

import (
	"log"

	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/repositories"
	"github.com/NSDN/nya-server/routes"
	"github.com/NSDN/nya-server/utils"
	"github.com/joho/godotenv"
)

func init() {
	// 设置日志输出前缀为日期及完整文件信息
	log.SetFlags(log.LstdFlags | log.Llongfile)

	// 从环境变量文件中加载环境变量
	err := godotenv.Load(
		configs.ENV_FILE,
		configs.ENV_FILE_DATABASE,
	)

	if err != nil {
		log.Fatal(utils.Messages.ENVIRONMENT_ERROR_NOT_FOUND)
	}
}

func main() {
	// 连接数据库
	// 使用 GORM 时不需要使用 `defer` 手动关闭数据库
	db := repositories.SetupDatabase()

	// 初始化路由
	r := routes.SetupRouter(db)

	// 从环境变量从获取端口号
	port := utils.GetENV(configs.ENV_APPLICATION_PORT)
	// 启动程序
	r.Run(":" + port)
}
