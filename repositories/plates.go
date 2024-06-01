package repositories

import (
	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/models"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
创建版块列表 - 数据库

	@param plates models/plates > Plate
	@returns _id[], err
*/
func InitPlateList(plates *[]models.Plate) ([]interface{}, error) {
	collection := getPlatesCollection()
	return insertManyToCollection(collection, plates)
}

// 获取版块列表 - 数据库
func GetPlateList() (*[]models.Plate, error) {
	// 从数据库中获取版块列表集合
	collection := getPlatesCollection()
	plates, err := findDataFromCollection(&[]models.Plate{}, collection, nil)

	if err != nil {
		return nil, err
	}

	return plates, nil
}

// 从数据库中获取版块列表集合
func getPlatesCollection() *mongo.Collection {
	return getCollection(configs.DB_COLLECTION_PLATES)
}
