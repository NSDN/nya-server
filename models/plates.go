package models

import "github.com/NSDN/nya-server/constants"

type Plate struct {
	// 版块 ID（兼路由名）
	ID string `json:"id" gorm:"column:id;size:100;primaryKey"`
	// 显示名
	Name string `json:"name" gorm:"column:name;size:100;not null"`
	// 版块说明
	Description string `json:"description" gorm:"column:description;type:text;not null"`
	// 背景图片（图床地址）
	Background string `json:"background" gorm:"column:background;type:text;not null"`
	// 画面类型（["commic", "article"]）
	PageType constants.PageType `json:"pageType" gorm:"column:page_type;size:20;not null"`
	// 排序锚点
	SortOrder int `json:"sortOrder" gorm:"column:sort_order;unique;not null;check:(sort_order >= 0)"`
}
