package repositories

import (
	"github.com/NSDN/nya-server/context"
	"github.com/NSDN/nya-server/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	context *context.AppContext
}

func NewUserRepository(
	context *context.AppContext,
) *UserRepository {
	return &UserRepository{context}
}

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
func (repository *UserRepository) InsertUser(user *models.User) (bool, error) {
	err := gorm.
		G[*models.User](repository.context.DB).
		Create(repository.context.DBContext, &user)

	return err == nil, err
}

func (repository *UserRepository) GetUser(
	username string,
) (*models.User, error) {
	return gorm.
		G[*models.User](repository.context.DB).
		Where("username = ?", username).
		Take(repository.context.DBContext)
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
