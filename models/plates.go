package models

import "github.com/NSDN/nya-server/constants"

type Plate struct {
	// 编号
	ID int `json:"id" bson:"id"`
	// 显示名
	Name string `json:"name" bson:"name"`
	// 路由名
	RouteName string `json:"routeName" bson:"routeName"`
	// 背景图片（图床地址）
	Background string `json:"background" bson:"background"`
	// 画面类型（["commic", "article"]）
	PageType constants.PageType `json:"pageType" bson:"pageType"`
}
