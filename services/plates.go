package services

import (
	"log"

	"github.com/NSDN/nya-server/constants"
	"github.com/NSDN/nya-server/models"
	"github.com/NSDN/nya-server/repositories"
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
			ID:         "localization",
			Name:       "喵玉汉化馆",
			Background: "https://i.imgur.com/ohQuzivl.jpg",
			PageType:   constants.COMMIC,
		},
		{
			ID:         "music",
			Name:       "喵玉咏唱组",
			Background: "https://i.imgur.com/IHo7tTyl.jpg",
			PageType:   constants.ARTICLE,
		},
		{
			ID:         "chat",
			Name:       "魔女的茶会",
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
