package models

// 用户
type User struct {
	// 用户编号
	UID string `json:"uid"`
	// 用户名
	Username string `json:"username"`
	// 昵称
	Nickname string `json:"nickname"`
	// 用户组
	UserGroup string `json:"userGruop"`
	// 头像
	Icon string `json:"icon"`
}
