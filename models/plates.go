package models

import "github.com/NSDN/nya-server/constants"

type Plate struct {
	// 版块 ID（兼路由名）
	ID string `json:"id" bson:"id"`
	// 显示名
	Name string `json:"name" bson:"name"`
	// 背景图片（图床地址）
	Background string `json:"background" bson:"background"`
	// 画面类型（["commic", "article"]）
	PageType constants.PageType `json:"pageType" bson:"pageType"`
}
