package models

// 用户一般信息
type User struct {
	// 用户编号
	UID string `json:"uid" bson:"uid"`
	// 用户名
	Username string `json:"username" bson:"username"`
	// 昵称
	Nickname string `json:"nickname" bson:"nickname"`
	// 用户组
	UserGroup string `json:"userGruop" bson:"userGruop"`
	// 头像
	Icon string `json:"icon" bson:"icon"`
}

// 用户完整信息
type UserFullInfo struct {
	// 用户编号
	UID string `json:"uid" bson:"uid,omitempty"`
	// 用户名
	Username string `json:"username" bson:"username,omitempty"`
	// 密码（明文哈希化后进行 Base64 编码后的值）
	Password string `json:"password" bson:"password,omitempty"`
	// 盐值（明文 Base64 编码后的值）
	Salt string `json:"salt" bson:"salt,omitempty"`
	// 昵称
	Nickname string `json:"nickname" bson:"nickname,omitempty"`
	// 用户组
	UserGroup string `json:"userGruop" bson:"userGruop,omitempty"`
	// 头像
	Icon string `json:"icon" bson:"icon,omitempty"`
}
