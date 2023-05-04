package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/NSDN/nya-server/utils"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *mongo.Client

// 连接数据库
//
// 需要在外部通过 defer 调用返回值以关闭连接
func SetupDatabase() *gorm.DB {
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
