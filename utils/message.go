package utils

import "fmt"

type MessageType struct {
	// 环境变量错误
	ENVIRONMENT_ERROR string

	// 连接成功
	CONNECT_SUCCESS string
}

// 信息类型
var messageType = MessageType{
	ENVIRONMENT_ERROR: "环境变量错误",

	CONNECT_SUCCESS: "连接成功",
}

type CustomMessages struct {
	// 找不到 .env 文件
	ENVIRONMENT_ERROR_NOT_FOUND string
	// 请在 .env 文件中设置名为 `<payload>` 的环境变量
	ENVIRONMENT_ERROR_NEED_ONE func(payload string) string

	// 成功连接至 <payload>
	CONNECT_SUCCESS func(payload string) string
}

// 自定义信息列表
var Messages = CustomMessages{
	ENVIRONMENT_ERROR_NOT_FOUND: fmt.Sprintf(
		"[%s] 找不到 .env 文件",
		messageType.ENVIRONMENT_ERROR,
	),

	ENVIRONMENT_ERROR_NEED_ONE: func(payload string) string {
		return fmt.Sprintf(
			"[%s] 请在 .env 文件中设置名为 `%s` 的环境变量",
			messageType.ENVIRONMENT_ERROR,
			payload,
		)
	},

	CONNECT_SUCCESS: func(payload string) string {
		return fmt.Sprintf(
			"[%s] 成功连接至%s！",
			messageType.CONNECT_SUCCESS,
			payload,
		)
	},
}
