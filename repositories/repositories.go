package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// 连接数据库
//
//	需要在外部通过 defer 调用返回值以关闭连接
func SetupDatabase() func() {
	// 从环境变量中获取数据库的连接信息
	uri := utils.GetENV(configs.ENV_MONGODB_URI)
	username := utils.GetENV(configs.ENV_MONGODB_USERNAME)
	password := utils.GetENV(configs.ENV_MONGODB_PASSWORD)

	// 创建 MongoDB 客户端选项
	clientOptions := options.Client().ApplyURI(uri).
		SetAuth(options.Credential{Username: username, Password: password})

	// 连接到 MongoDB
	// context.Background() 表示一个空的 goroutine 上下文
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	Client = client

	// 结束时关闭连接的方法
	disconnect := func() {
		err = client.Disconnect(context.Background())

		if err != nil {
			log.Fatal(err)
		}
	}

	// 检查连接是否成功
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
		disconnect()
	}

	fmt.Println(utils.Messages.CONNECT_SUCCESS(" MongoDB"))

	return disconnect
}
