package models

// 注册信息
type RegisterInfo struct {
	// 用户名
	Username string `json:"username" bson:"username"`
	// 密码
	Password string `json:"password" bson:"password"`
	// 再次确认密码
	ConfirmPassword string `json:"confirmPassword" bson:"confrimPassword"`
}

// 登入信息
type LoginInfo struct {
	// 用户名
	Username string `json:"username" bson:"username"`
	// 密码
	Password string `json:"password" bson:"password"`
}
