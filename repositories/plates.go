package repositories

import (
	"context"

	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
创建版块列表 - 数据库

	@param plates models/plates > Plate
	@returns _id[], err
*/
func InitPlateList(plates []interface{}) ([]interface{}, error) {
	// 从数据库中获取版块列表集合
	collection := getPlatesCollection()

	// 创建上下文
	c := context.Background()
	// 插入版块列表
	result, err := collection.InsertMany(c, plates)

	return result.InsertedIDs, err
}

// 获取版块列表 - 数据库
func GetPlateList() ([]models.Plate, error) {
	// 从数据库中获取版块列表集合
	collection := getPlatesCollection()

	// 创建上下文
	c := context.Background()

	// 获取查询版块列表集合的游标
	cursor, err := collection.Find(c, bson.D{})

	if err != nil {
		return nil, err
	}

	// 函数结束时关闭游标
	defer cursor.Close(c)

	var plates []models.Plate

	// 解码游标结果集至 plates 结构体
	err = cursor.All(c, &plates)

	if err != nil {
		return nil, err
	}

	return plates, nil
}

// 从数据库中获取版块列表集合
func getPlatesCollection() *mongo.Collection {
	return Client.
		Database(configs.DATABASE_NAME).
		Collection(configs.DB_COLLECTION_PLATES)
}
