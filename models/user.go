package models

// 用户一般信息
type UserPublicInfo struct {
	// 用户编号
	ID int64 `json:"id" gorm:"column:id;primaryKey"`
	// 用户名
	Username string `json:"username"`
	// 昵称
	Nickname string `json:"nickname"`
	// 用户组
	UserGroup string `json:"userGruop"`
	// 头像
	Icon string `json:"icon"`
}

// 用户完整信息
type User struct {
	// 用户编号，数据库自动生成
	UID string `json:"uid" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	// 用户名
	Username string `json:"username" gorm:"unique;not null"`
	// 密码（明文哈希化后进行 Base64 编码后的值）
	Password string `json:"password" gorm:"not null"`
	// 盐值（明文 Base64 编码后的值）
	Salt string `json:"salt" gorm:"not null"`
	// 昵称
	Nickname string `json:"nickname" gorm:"not null"`
	// 用户组
	UserGroup string `json:"userGroup" gorm:"not null"`
	// 头像
	Icon string `json:"icon"`
}
