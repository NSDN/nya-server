package services

import (
	"github.com/NSDN/nya-server/models"
	"github.com/NSDN/nya-server/repositories"
	"github.com/NSDN/nya-server/utils"
)

// 获取帖文列表 - 服务
//
// 从数据库中找出所有文章列表，然后根据传入的版块路由名来过筛选出当前版块的文章列表。
func GetArticles(plate string) (*[]models.Article, error) {
	articles, err := repositories.GetArticles()

	if err != nil {
		return nil, err
	}

	articles = utils.FilterSlice(
		articles,

		func(item *models.Article, index int) bool {
			return item.Common.Plate == plate
		},
	)

	return articles, nil
}

// 创建文章 - 服务
func CreateArticle(article *models.Article) (bool, error) {
	return repositories.CreateArticle(article)
}
