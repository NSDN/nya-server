package plates

type Plate struct {
	ID int `json:"id"`
	// 显示名
	Name string `json:"name"`
	// 路由名
	RouteName string `json:"routeName"`
	// 版块名颜色（同 css）
	NameColor string `json:"nameColor"`
	// 背景图片（图床地址）
	Background string `json:"background"`
}
