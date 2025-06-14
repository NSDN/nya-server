package main

import (
	"fmt"
	"log"

	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/context"
	"github.com/NSDN/nya-server/routes"
	"github.com/NSDN/nya-server/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

// 连接数据库
//
// 需要在外部通过 defer 调用返回值以关闭连接
func initDatabase() *gorm.DB {
	// 从环境变量中获取数据库的连接信息
	connectString := fmt.Sprintf(
		"dbname=%s user=%s password=%s sslmode=disable",
		utils.GetENV("POSTGRES_DB"),
		utils.GetENV("POSTGRES_USER"),
		utils.GetENV("POSTGRES_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{})

	if err != nil {
		log.Fatal("Database open error: ", err)
	}

	sqlDB, err := db.DB()
	err = sqlDB.Ping()

	if err != nil {
		log.Fatal("Ping error: ", err)
	}

	return db
}

func main() {
	// 连接数据库
	// 使用 GORM 时不需要使用 `defer` 手动关闭数据库
	db := initDatabase()

	appContext := &context.AppContext{
		DB: db,
	}

	// 初始化路由
	r := routes.SetupRouter(appContext)

	// 从环境变量从获取端口号
	port := utils.GetENV(configs.ENV_APPLICATION_PORT)
	// 启动程序
	r.Run(":" + port)
}
