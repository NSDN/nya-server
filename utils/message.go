package utils

import "fmt"

type MessageType struct {
	// 环境变量错误
	ENVIRONMENT_ERROR string
	// 连接成功
	CONNECT_SUCCESS string
	// 认证失败
	AUTHORIZE_FAILED string
}

// 信息类型
var messageType = MessageType{
	ENVIRONMENT_ERROR: "环境变量错误",
	CONNECT_SUCCESS:   "连接成功",
	AUTHORIZE_FAILED:  "认证失败",
}

type CustomMessages struct {
	// 找不到 .env 文件
	ENVIRONMENT_ERROR_NOT_FOUND string
	// 请在 .env 文件中设置名为 `<payload>` 的环境变量
	ENVIRONMENT_ERROR_NEED_ONE func(payload string) string

	// 成功连接至 <payload>
	CONNECT_SUCCESS func(payload string) string

	// 缺少 <payload> 参数
	AUTHORIZE_FAILED_MISSING_PARAMETER func(payload string) string
	// 验证信息格式不正确
	AUTHORIZE_FAILED_BAD_PARAMETER string
	// 参数 <payload> 错误
	AUTHORIZE_FAILED_WRONG_PARAMETER func(payload string) string
	// 找不到用户
	AUTHORIZE_FAILED_NO_USER string
	// 密码不一致
	AUTHORIZE_FAILED_WRONG_PASSWORD string
	// 错误的令牌
	AUTHORIZE_FAILED_WRONG_TOKEN string
}

// 自定义信息列表
var Messages = CustomMessages{
	ENVIRONMENT_ERROR_NOT_FOUND: fmt.Sprintf(
		"[%s] 找不到 .env 文件。",
		messageType.ENVIRONMENT_ERROR,
	),

	ENVIRONMENT_ERROR_NEED_ONE: func(payload string) string {
		return fmt.Sprintf(
			"[%s] 请在 .env 文件中设置名为 `%s` 的环境变量。",
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

	AUTHORIZE_FAILED_MISSING_PARAMETER: func(payload string) string {
		return fmt.Sprintf(
			"[%s] 缺少%s参数。",
			messageType.AUTHORIZE_FAILED,
			payload,
		)
	},

	AUTHORIZE_FAILED_BAD_PARAMETER: fmt.Sprintf(
		"[%s] 验证信息格式不正确",
		messageType.AUTHORIZE_FAILED,
	),

	AUTHORIZE_FAILED_WRONG_PARAMETER: func(payload string) string {
		return fmt.Sprintf(
			"[%s] 参数%s错误。",
			messageType.AUTHORIZE_FAILED,
			payload,
		)
	},

	AUTHORIZE_FAILED_NO_USER: fmt.Sprintf(
		"[%s] 找不到用户。",
		messageType.AUTHORIZE_FAILED,
	),

	AUTHORIZE_FAILED_WRONG_PASSWORD: fmt.Sprintf(
		"[%s] 密码不一致。",
		messageType.AUTHORIZE_FAILED,
	),

	AUTHORIZE_FAILED_WRONG_TOKEN: fmt.Sprintf(
		"[%s] 错误的令牌。",
		messageType.AUTHORIZE_FAILED,
	),
}
