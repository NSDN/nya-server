package configs

// 环境变量
const (
	// 环境变量文件名
	ENV_FILE = ".env"
	// MongoDB 环境变量文件名
	ENV_FILE_MONGODB = ".env.mongodb"

	// 程序启动端口
	ENV_APPLICATION_PORT = "APPLICATION_PORT"

	// 数据库 URI
	ENV_MONGODB_URI = "MONGODB_URI"
	// 数据库用户名（不要共享到 GitHub 等公开环境）
	ENV_MONGODB_USERNAME = "MONGO_INITDB_ROOT_USERNAME"
	// 数据库密码（不要共享到 GitHub 等公开环境）
	ENV_MONGODB_PASSWORD = "MONGO_INITDB_ROOT_PASSWORD"
)

// 数据库
const (
	// 数据库名称
	DATABASE_NAME = "miaoYuDian"
	// 版块列表集合
	DB_COLLECTION_PLATES = "plates"
)

// 认证授权
const (
	// 密码哈希计算的成本，越高越安全但计算所需时间和资源越多，取值范围是 4 - 31。
	BCRYPT_COST = 10
	// 盐值长度
	SALT_LENGTH = 16
)
