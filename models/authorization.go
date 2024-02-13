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

// 用户验证信息
type UserAuthorizateInfo struct {
	// 用户编号
	UID string `json:"uid" bson:"uid"`
	// 用户名
	Username string `json:"username" bson:"username"`
	// 密码（明文哈希化后进行 Base64 编码后的值）
	Password string `json:"password" bson:"password"`
	// 盐值（明文 Base64 编码后的值）
	Salt string `json:"salt" bson:"salt"`
}
