package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewTopicRequestData struct {
	// 作者
	Author User `json:"author" bson:"author"`
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

// 帖子表
type Topic struct {
	// 帖子 ID
	ID int64 `gorm:"primaryKey;autoIncrement"`
	// 作者 ID（User 外键：即不是 Topic 自身的属性，为了关联 User 而存在。）
	AuthorID int64 `gorm:"not null"`
	// 作者实例（GORM 预加载用，如果外键引用的 User 中的属性不是 `ID`，需要使用 `references` 手动指定。）
	Author User `gorm:"foreignKey:AuthorID;"`
	// 版块（版块ID）
	PlateID string `gorm:"not null"`
	// 版块实例
	Plate Plate `gorm:"foreignKey:PlateID"`
	// 标题
	Title string `gorm:"not null"`
	// 帖子类型
	TopicType string `gorm:"size:20;not null"`
	// 预览图链接
	ThumbnailLink string `gorm:"not null"`
	// TAG
	Tag []string `gorm:"type:varchar(20)[]"`
	// 创建时间
	CreatedAt time.Time `gorm:"not null"`
	// 更新时间
	UpdatedAt time.Time `gorm:"not null"`
}

type TopicWidthID struct {
	// 帖文ID（由数据库自增生成）
	TopicID primitive.ObjectID `json:"topicID" bson:"_id"`
	// 作者
	Author User `json:"author" bson:"author"`
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
	List []FloorItem `json:"list" bson:"list"`
}

type FloorItem struct {
	// 层数
	Level int `json:"level" bson:"level"`
	// 作者
	Author User `json:"author" bson:"author"`
	// 创建日
	CreationDate time.Time `json:"creationDate" bson:"creationDate"`
	// 回帖文章类型
	BodyType string `json:"bodyType" bason:"bodyType"`
	// 回帖正文
	Body string `json:"body" bson:"body"`
}
