package services

import (
	"log"

	"github.com/NSDN/nya-server/constants"
	"github.com/NSDN/nya-server/context"
	"github.com/NSDN/nya-server/models"
	"github.com/NSDN/nya-server/repositories"
)

type PlateService struct {
	context    *context.AppContext
	repository *repositories.PlateRepository
}

func NewPlateService(context *context.AppContext) *PlateService {
	return &PlateService{
		context:    context,
		repository: repositories.NewPlateRepository(context),
	}
}

// 创建版块列表 - 服务
func (service *PlateService) InitPlateList() {
	// 从数据库中获取版块列表
	plates, err := service.GetPlates()

	if err != nil {
		log.Fatal(err)
	}

	// 如果有既存的列表，则不做任何操作
	if len(plates) > 0 {
		return
	}

	// 如果没有既存列表，则插入列表
	newPlates := []models.Plate{
		{
			ID:          "localization",
			Name:        "喵玉汉化馆",
			Description: "",
			Background:  "https://i.imgur.com/ohQuzivl.jpg",
			PageType:    constants.COMIC,
			SortOrder:   0,
		},
		{
			ID:          "music",
			Name:        "喵玉咏唱组",
			Description: "",
			Background:  "https://i.imgur.com/IHo7tTyl.jpg",
			PageType:    constants.ARTICLE,
			SortOrder:   1,
		},
		{
			ID:          "chat",
			Name:        "魔女的茶会",
			Description: "",
			Background:  "https://i.imgur.com/JsWkJ4jl.jpg",
			PageType:    constants.ARTICLE,
			SortOrder:   2,
		},
	}

	_, err = service.repository.InitPlateList(newPlates)

	if err != nil {
		log.Fatal(err)
	}
}

// 获取版块列表 - 服务
func (service *PlateService) GetPlates() ([]models.Plate, error) {
	return service.repository.GetPlates()
}
