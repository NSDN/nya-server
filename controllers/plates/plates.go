package plates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlateList struct {
	ID int `json:"id"`
	// 版块名
	Title string `json:"title"`
	// 版块名颜色（同 css）
	TitleColor string `json:"titleColor"`
	// 背景图片（图床地址）
	Background string `json:"background"`
	// 路由名
	RouteName string `json:"routeName"`
}

func GetPlateList(context *gin.Context) {

	list := []PlateList{
		{
			ID:         0,
			Title:      "汉化区",
			TitleColor: "#fff",
			Background: "https://static-event.benghuai.com/new_mihoyo_homepage/images/download/cg/origin/2020-10-22.jpg",
			RouteName:  "TranslatePlate",
		},
		{
			ID:        1,
			Title:     "sample01",
			RouteName: "Sample01",
		},
		{
			ID:        2,
			Title:     "sample02",
			RouteName: "Sample02",
		},
	}

	context.JSON(http.StatusOK, list)
}
