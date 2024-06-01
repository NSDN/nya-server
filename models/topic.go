package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewTopicRequestData struct {
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
	// 文章类型
	BodyType string `json:"bodyType" bason:"bodyType"`
	// 文章正文
	Body string `json:"body" bson:"body"`
}

// 此处不使用 ID 是为了避免插入时覆盖 mongodb 自增ID
type Topic struct {
	// 作者
	Author string `json:"author" bson:"author"`
	// 版块（版块ID）
	Plate string `json:"plate" bson:"plate"`
	// 标题
	Title string `json:"title" bson:"title"`
	// TAG
	Tag []string `json:"tag" bson:"tag"`
	// 创建日
	CreationDate time.Time `json:"creationDate" bson:"creationDate"`
}

type TopicWidthID struct {
	// 帖文ID（由数据库自增生成）
	TopicID primitive.ObjectID `json:"topicID" bson:"_id"`
	// 作者
	Author string `json:"author" bson:"author"`
	// 版块（版块ID）
	Plate string `json:"plate" bson:"plate"`
	// 标题
	Title string `json:"title" bson:"title"`
	// TAG
	Tag []string `json:"tag" bson:"tag"`
	// 创建日
	CreationDate time.Time `json:"creationDate" bson:"creationDate"`
}

type TopicFloors struct {
	// 帖文ID（由 topics 数据库自增生成后手动添加至 floors 数据库里）
	TopicID primitive.ObjectID `json:"topicID" bson:"_id"`
	// 楼层列表
	List []FloorItem
}

type FloorItem struct {
	// 层数
	Level int `json:"level" bson:"level"`
	// 作者
	Author string `json:"author" bson:"author"`
	// 创建日
	CreationDate time.Time `json:"creationDate" bson:"creationDate"`
	// 回帖文章类型
	BodyType string `json:"bodyType" bason:"bodyType"`
	// 回帖正文
	Body string `json:"body" bson:"body"`
}
