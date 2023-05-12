package plates

type PageType string

const (
	COMMIC  PageType = "commic"
	ARTICLE PageType = "article"
)

type Plate struct {
	// 编号
	ID int `json:"id"`
	// 显示名
	Name string `json:"name"`
	// 路由名
	RouteName string `json:"routeName" bson:"routeName"`
	// 背景图片（图床地址）
	Background string `json:"background"`
	// 画面类型（["commic", "article"]）
	PageType PageType `json:"pageType" bson:"pageType"`
}
