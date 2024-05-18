package models

// 文章
type Article struct {
	// 文章标题
	Title string `json:"title" bson:"title"`
	// 文章所属版块（版块路由名）
	Plate string `json:"plate" bson:"plate"`
}
