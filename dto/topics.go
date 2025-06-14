package dto

import "time"

type TopicListItemDTO struct {
	// 帖子 ID（经过编码的）
	ID string `json:"id"`
	// 标题
	Title string `json:"title"`
	// 预览图链接
	ThumbnailLink string `json:"thumbnailLink"`
	// 更新时间
	UpdatedAt time.Time `json:"updatedAt"`
}
