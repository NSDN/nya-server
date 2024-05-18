package services

import (
	"log"

	"github.com/NSDN/nya-server/constants"
	"github.com/NSDN/nya-server/models"
	"github.com/NSDN/nya-server/repositories"
	"github.com/NSDN/nya-server/utils"
)

// 创建版块列表 - 服务
func InitPlateList() {
	// 从数据库中获取版块列表
	plates, err := GetPlateList()

	if err != nil {
		log.Fatal(err)
	}

	// 如果有既存的列表，则不做任何操作
	if len(*plates) > 0 {
		return
	}

	// 如果没有既存列表，则插入列表
	newPlates := []models.Plate{
		{
			ID:         0,
			Name:       "喵玉汉化馆",
			RouteName:  "localization",
			Background: "https://i.imgur.com/ohQuzivl.jpg",
			PageType:   constants.COMMIC,
		},
		{
			ID:         1,
			Name:       "喵玉咏唱组",
			RouteName:  "music",
			Background: "https://i.imgur.com/IHo7tTyl.jpg",
			PageType:   constants.ARTICLE,
		},
		{
			ID:         2,
			Name:       "魔女的茶会",
			RouteName:  "chat",
			Background: "https://i.imgur.com/JsWkJ4jl.jpg",
			PageType:   constants.ARTICLE,
		},
	}

	_, err = repositories.InitPlateList(&newPlates)

	if err != nil {
		log.Fatal(err)
	}
}

// 获取版块列表 - 服务
func GetPlateList() (*[]models.Plate, error) {
	return repositories.GetPlateList()
}

// TODO: 等创建文章功能完成后应删除此函数
//
// 创建文章列表假数据 - 服务
func InitDummyArticles() {
	articles, err := repositories.GetArticles()

	if err != nil {
		log.Fatal(err)
	}

	if len(*articles) > 0 {
		return
	}

	articles = &[]models.Article{
		{
			Title: "青年在选择职业时的考虑",
			Plate: "chat",
		},
		{
			Title: "德意志意识形态",
			Plate: "chat",
		},
		{
			Title: "1844年经济学哲学手稿",
			Plate: "music",
		},
	}

	_, err = repositories.InitDummyArticles(articles)

	if err != nil {
		log.Fatal(err)
	}
}

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
			return item.Plate == plate
		},
	)

	return articles, nil
}
