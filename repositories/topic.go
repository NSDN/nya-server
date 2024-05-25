package repositories

import (
	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// 获取文章集合
func getArticleCollection() *mongo.Collection {
	return getCollection(configs.DB_COLLECTION_ARTICLES)
}

// 获取帖文列表 - 数据库
func GetArticles() (*[]models.Article, error) {
	collection := getArticleCollection()
	articles, err := findDataFromCollection(&[]models.Article{}, collection)

	if err != nil {
		return nil, err
	}

	return articles, nil
}

// 创建文章 - 数据库
func CreateArticle(article *models.Article) (bool, error) {
	collection := getArticleCollection()
	return insertOneToCollection(collection, article)
}
