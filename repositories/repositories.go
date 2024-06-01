package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/utils"
	"go.mongodb.org/mongo-driver/bson"
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

// 插入列表数据进集合
//
// struct/any 皆不能作为 interface{} 类型传递，因此需要生成一遍 interface{} 类型的数据。
func insertManyToCollection[T any](
	collection *mongo.Collection,
	data *[]T,
) ([]interface{}, error) {
	var payload []interface{}

	for _, item := range *data {
		payload = append(payload, item)
	}

	// 创建上下文
	c := context.Background()
	// 插入数据
	result, err := collection.InsertMany(c, payload)
	return result.InsertedIDs, err
}

// 插入单条数据进集合
func insertOneToCollection[T any](
	collection *mongo.Collection,
	data *T,
) (*mongo.InsertOneResult, error) {
	insertResult, err := collection.InsertOne(context.TODO(), data)

	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

// 从数据库中取出集合
//
//	collection 集合名
func getCollection(collection string) *mongo.Collection {
	return Client.
		Database(configs.DATABASE_NAME).
		Collection(collection)
}

// 从数据库的集合中找出单个数据
func findOneDataFromCollection[T any](
	model *T,
	collection *mongo.Collection,
	filter interface{},
) (*T, error) {
	// 创建上下文
	c := context.Background()

	if filter == nil {
		filter = bson.D{}
	}

	// 获取集合的游标
	err := collection.FindOne(c, filter).Decode(model)

	if err != nil {
		return nil, err
	}

	return model, nil
}

// 从数据库的集合中找出数据集合
func findDataFromCollection[T any](
	model *T,
	collection *mongo.Collection,
	filter interface{},
) (*T, error) {
	// 创建上下文
	c := context.Background()

	if filter == nil {
		filter = bson.D{}
	}

	// 获取集合的游标
	cursor, err := collection.Find(c, filter)

	if err != nil {
		return nil, err
	}

	// 函数结束时关闭游标
	defer cursor.Close(c)

	// 解码游标结果集至结构体
	err = cursor.All(c, model)

	if err != nil {
		return nil, err
	}

	return model, nil
}
