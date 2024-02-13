package repositories

import (
	"context"

	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 从数据库中获取用户列表集合
func getUsersCollection() *mongo.Collection {
	return Client.
		Database(configs.DATABASE_NAME).
		Collection(configs.DB_COLLECTION_USERS)
}

// 获取用户总数
func GetUserCount() (int64, error) {
	collection := getUsersCollection()
	return collection.CountDocuments(context.TODO(), bson.D{})
}

// 插入新用户
func InsertNewUser(newUser *models.UserFullInfo) (bool, error) {
	collection := getUsersCollection()
	_, err := collection.InsertOne(context.TODO(), newUser)

	if err != nil {
		return false, err
	}

	return true, nil
}
