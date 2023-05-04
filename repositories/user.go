package repositories

import (
	"github.com/NSDN/nya-server/models"
)

// 从数据库中获取用户列表集合
// func getUsersCollection() *mongo.Collection {
func getUsersCollection() {
	// return getCollection(configs.DB_COLLECTION_USERS)
}

// 获取用户总数
func GetUserCount() (int64, error) {
	// collection := getUsersCollection()
	// return collection.CountDocuments(context.TODO(), bson.D{})
	return 0, nil
}

// 插入新用户
func InsertNewUser(newUser *models.UserFullInfo) (bool, error) {
	// collection := getUsersCollection()
	// _, err := insertOneToCollection(collection, newUser)

	// if err != nil {
	// 	return false, err
	// }

	// return true, nil
	return false, nil
}

// 获取用户列表 - 数据库
// func GetUserList() (*[]models.User, error) {
func GetUserList() error {
	// collection := getUsersCollection()

	// users, err := findDataFromCollection(&[]models.User{}, collection, nil)

	// if err != nil {
	// 	return nil, err
	// }

	// return users, nil
	return nil
}
