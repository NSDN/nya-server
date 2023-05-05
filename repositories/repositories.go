package repositories

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/NSDN/nya-server/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDatabase() {
	MONGODB_URI := "MONGODB_URI"
	uri := os.Getenv(MONGODB_URI)

	if uri == "" {
		log.Fatal(utils.Messages.ENVIRONMENT_ERROR_NEED_ONE(MONGODB_URI))
	}

	// 创建 MongoDB 客户端选项
	clientOptions := options.Client().ApplyURI(uri)

	// 连接到 MongoDB
	// context.Background() 表示一个空的 goroutine 上下文
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// 检查连接是否成功
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(utils.Messages.CONNECT_SUCCESS(" MongoDB"))
}
