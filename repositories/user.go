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
	return getCollection(configs.DB_COLLECTION_USERS)
}

// 获取用户总数
func GetUserCount() (int64, error) {
	collection := getUsersCollection()
	return collection.CountDocuments(context.TODO(), bson.D{})
}

// 插入新用户
func InsertNewUser(newUser *models.UserFullInfo) (bool, error) {
	collection := getUsersCollection()
	return insertOneToCollection(collection, newUser)
}

// 获取用户列表 - 数据库
func GetUserList() (*[]models.User, error) {
	collection := getUsersCollection()

	users, err := findDataFromCollection(&[]models.User{}, collection)

	if err != nil {
		return nil, err
	}

	return users, nil
}
