package models

// 文章
type Article struct {
	// 帖子共通信息
	Common TopicCommon `json:"common" bason:"common"`
	// 楼层
	Floors []Floor `json:"floors" bason:"floors"`
}
