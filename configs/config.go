package configs

// 环境变量名
const (
	// 环境变量文件名
	ENV_FILE = ".env"
	// 数据库环境变量文件名
	ENV_FILE_DATABASE = ".env.postgres"

	// 程序启动端口
	ENV_APPLICATION_PORT = "APPLICATION_PORT"
	// 令牌签名用键
	ENV_TOKEN_KEY = "TOKEN_KEY"
)

// 数据库
const (
	// 数据库名称
	DATABASE_NAME = "nyasama"
	// 用户列表集合
	DB_COLLECTION_USERS = "users"
	// 版块列表集合
	DB_COLLECTION_PLATES = "plates"
	// 帖子列表集合
	DB_COLLECTION_TOPICS = "topics"
	// 帖子楼层列表集合
	DB_COLLECTION_TOPIC_FLOORS = "topicFloors"
)

// 认证授权
const (
	// 密码哈希计算的成本，越高越安全但计算所需时间和资源越多，取值范围是 4 - 31。
	BCRYPT_COST = 10
	// 盐值长度
	SALT_LENGTH = 16
	// 密码最长下标
	// 加密算法限制明文密码长度小于 72，而盐值占据 16，故密码下标设为最大 55
	PASSWORD_MAX_INDEX = 55
	// Authorization 认证方式
	AUTHENTICATION_TYPE = "Bearer "
)

// 上下文键名
const (
	// 声明集
	CONTEXT_KEY_CLAIMS = "claims"
)

// 请求头
const (
	// 认证信息
	HTTP_HEADER_AUTHORIZATION = "Authorization"
)
