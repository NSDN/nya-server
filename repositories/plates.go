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
	return insertManyDataToCollection(collection, plates)
}

// 获取版块列表 - 数据库
func GetPlateList() (*[]models.Plate, error) {
	// 从数据库中获取版块列表集合
	collection := getPlatesCollection()
	plates, err := findDataFromCollection(&[]models.Plate{}, collection)

	if err != nil {
		return nil, err
	}

	return plates, nil
}

// 从数据库中获取版块列表集合
func getPlatesCollection() *mongo.Collection {
	return getCollection(configs.DB_COLLECTION_PLATES)
}

// TODO: 等创建文章功能完成后应删除此函数
//
// 生成假帖文列表 - 数据库
func InitDummyArticles(articles *[]models.Article) ([]interface{}, error) {
	collection := getCollection(configs.DB_COLLECTION_ARTICLES)
	return insertManyDataToCollection(collection, articles)
}

// 获取帖文列表 - 数据库
func GetArticles() (*[]models.Article, error) {
	collection := getCollection(configs.DB_COLLECTION_ARTICLES)
	articles, err := findDataFromCollection(&[]models.Article{}, collection)

	if err != nil {
		return nil, err
	}

	return articles, nil
}
