package models

// 登入信息
type LoginInfo struct {
	// 用户名
	Username string `json:"username" bson:"username"`
	// 密码
	Password string `json:"password" bson:"password"`
}
