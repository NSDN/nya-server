package plates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlateList struct {
	ID int `json:"id"`
	// 显示名
	Name string `json:"name"`
	// 路由名
	RouteName string `json:"routeName"`
	// 版块名颜色（同 css）
	TitleColor string `json:"titleColor"`
	// 背景图片（图床地址）
	Background string `json:"background"`
}

func GetPlateList(context *gin.Context) {
	list := []PlateList{
		{
			ID:         0,
			Name:       "汉化区",
			RouteName:  "LocalizationPlate",
			TitleColor: "#fff",
			Background: "https://static-event.benghuai.com/new_mihoyo_homepage/images/download/cg/origin/2020-10-22.jpg",
		},
		{
			ID:        1,
			Name:      "sample01",
			RouteName: "Sample01",
		},
		{
			ID:        2,
			Name:      "sample02",
			RouteName: "Sample02",
		},
	}

	context.JSON(http.StatusOK, list)
}
