package models

import "time"

type TopicCommon struct {
	// 作者
	Author string `json:"author" bson:"author"`
	// 版块（版块 ID）
	Plate string `json:"plate" bson:"plate"`
	// 标题
	Title string `json:"title" bson:"title"`
	// TAG
	Tag []string `json:"tag" bson:"tag"`
	// 创建日
	CreationDate time.Time `json:"creationDate" bson:"creationDate"`
}

type Floor struct {
	// 层数
	Level int `json:"level" bson:"level"`
	// 作者
	Author string `json:"author" bson:"author"`
	// 帖文类型
	BodyType string `json:"bodyType" bson:"bodyType"`
	// 帖文
	Body string `json:"body" bson:"body"`
}
