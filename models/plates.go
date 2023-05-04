package models

import "github.com/NSDN/nya-server/constants"

type Plate struct {
	// 版块 ID（兼路由名）
	ID string `json:"id" gorm:"size:100;primaryKey"`
	// 显示名
	Name string `json:"name" gorm:"size:100;not null"`
	// 版块说明
	Description string `json:"description" gorm:"type:text"`
	// 背景图片（图床地址）
	Background string `json:"background" gorm:"type:text"`
	// 画面类型（["commic", "article"]）
	PageType constants.PageType `json:"pageType" gorm:"size:20;not null;"`
}
