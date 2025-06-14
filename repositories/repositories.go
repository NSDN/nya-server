package repositories

import (
	"context"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

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
