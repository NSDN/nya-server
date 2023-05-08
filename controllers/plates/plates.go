package plates

import (
	"net/http"

	platesModels "github.com/NSDN/nya-server/models/plates"
	"github.com/gin-gonic/gin"
)

func GetPlateList(context *gin.Context) {
	list := []platesModels.Plate{
		{
			ID:         0,
			Name:       "喵玉汉化馆",
			RouteName:  "localization",
			NameColor:  "#fff",
			Background: "https://static-event.benghuai.com/new_mihoyo_homepage/images/download/cg/origin/2020-10-22.jpg",
			PageType:   platesModels.COMMIC,
		},
		{
			ID:        1,
			Name:      "喵玉咏唱组",
			RouteName: "music",
			PageType:  platesModels.ARTICLE,
		},
		{
			ID:        2,
			Name:      "魔女的茶会",
			RouteName: "chat",
			PageType:  platesModels.ARTICLE,
		},
	}

	context.JSON(http.StatusOK, list)
}
