package plates

import (
	"log"

	"github.com/NSDN/nya-server/constants"
	models "github.com/NSDN/nya-server/models/plates"
	repositories "github.com/NSDN/nya-server/repositories/plates"
)

// 创建版块列表 - 服务
func InitPlateList() {
	// 从数据库中获取版块列表
	plates, err := GetPlateList()

	if err != nil {
		log.Fatal(err)
	}

	// 如果有既存的列表，则不做任何操作
	if len(plates) > 0 {
		return
	}

	// 如果没有既存列表，则插入列表
	newPlates := []interface{}{
		models.Plate{
			ID:         0,
			Name:       "喵玉汉化馆",
			RouteName:  "localization",
			Background: "https://i.imgur.com/ohQuzivl.jpg",
			PageType:   constants.COMMIC,
		},

		models.Plate{
			ID:         1,
			Name:       "喵玉咏唱组",
			RouteName:  "music",
			Background: "https://i.imgur.com/IHo7tTyl.jpg",
			PageType:   constants.ARTICLE,
		},

		models.Plate{
			ID:         2,
			Name:       "魔女的茶会",
			RouteName:  "chat",
			Background: "https://i.imgur.com/JsWkJ4jl.jpg",
			PageType:   constants.ARTICLE,
		},
	}

	_, err = repositories.InitPlateList(newPlates)

	if err != nil {
		log.Fatal(err)
	}
}

// 获取版块列表 - 服务
func GetPlateList() ([]models.Plate, error) {
	return repositories.GetPlateList()
}
