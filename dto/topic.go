package dto

type TopicResponse struct {
	// 帖子 ID（经过编码的）
	ID string `json:"id"`
	// 用户
	Author string `json:"author"`
}
