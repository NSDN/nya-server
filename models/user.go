package models

// 用户
type User struct {
	// 用户编号
	UID string
	// 用户名
	Username string
	// 密码（明文哈希化后进行 Base64 编码后的值）
	Password string
	// 盐值（明文 Base64 编码后的值）
	Salt string
}
