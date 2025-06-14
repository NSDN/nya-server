# 喵玉殿论坛服务端

## 开发前需要安装的工具

- [Golang]\: 编程语言（自带包管理功能）。
- [golang-migrate/migrate]\: 数据库迁移工具（可通过 `go install` 的方式安装）。
- [Docker]（可选）: 容器化平台（用来干净地安装数据库）。

二者的版本请见 [go.mod](./go.mod)。

## 代码获取方式

```bash
git clone https://github.com/NSDN/nya-server.git
```

## 参与开发请看

参与开发请务必查看 [CONTRIBUTING.md](./CONTRIBUTING.md) 文件。

## 启动服务端的方式

**\* 所有步骤必须严格按顺序执行。**

### 一、创建环境变量

- 基于所有的 `.env.*.example` 文件创建自己的环境变量文件（注意这是敏感数据请勿使用 Git 管理）。
- 如果不了解示例文件中的字段的作用，可以仅复制示例文件并删除 `.example` 的部分。

### 二、启动容器以使用数据库

**\* 这一步需要事先安装了 [Docker]。不采用 [Docker] 方式请自行启动数据库并跳过这一步 。**

1. 执行 [Docker] 的可执行文件。

2. 执行 [Docker Compose](./docker-compose.yml) 文件（这一步会自动安装数据库）：

   ```bash
   docker compose up --detach
   ```

_其他相关命令：_

- 查看容器的日志：

  ```bash
  docker compose logs nya-db
  ```

- 进入容器：

  ```bash
  docker compose exec -it nya-db bash
  ```

- 退出并销毁容器：

  ```bash
  docker compose down
  ```

### 三、迁移数据库

**\* 这一步需要事先安装了 [golang-migrate/migrate]。**

```bash
# up [N]       Apply all or N up migrations
migrate -source file://./database/migrations -database "postgres://{{user}}:{{password}}@localhost:5432/forum?sslmode=disable" up
```

- `user` 和 `password` 使用定义在 [`.env.postgres`](./.env.postgres.example) 中的值。
- `host`、`port`、`dbname` 等亦可自定义，但需自行在环境变量文件中增加字段。

_其他相关命令：_

- 撤销迁移：

  ```bash
  # down [N] [-all]    Apply all or N down migrations
  migrate -source file://./database/migrations -database "postgres://{{user}}:{{password}}@localhost:5432/forum?sslmode=disable" down
  ```

- 创建迁移文件：

  ```bash
  migrate create -dir database/migrations -ext sql -seq {{migration_name}}
  ```

  - `-ext`: 文件扩展名。
  - `-seq`: 迁移文件名称。

- 删除迁移文件：

  ```bash
  rm database/migrations/{{migration_name}}.sql
  ```

### 四、启动 Golang 程序

\* 这一步需要事先安装了 [Golang]。

1. 初始化依赖：

   ```bash
   go mod tidy
   ```

2. 校验依赖：

   ```bash
   go mod verify
   ```

3. 启动程序：

   ```bash
   go run .
   ```

### 五、验证项目已正常启动

```shell
# 返回 `pong` 则表示启动成功
curl -X GET 'http://localhost:10127/ping'
```

## 依赖库

版本请查看 [go.mod](./go.mod)。

| 名称                     | 功能                      |
| ------------------------ | ------------------------- |
| [Golang]                 | 编程语言                  |
| [Gin]                    | Web 服务端框架            |
| [GORM]                   | ORM                       |
| [PostgreSQL]             | 数据库                    |
| [golang-migrate/migrate] | [Golang] 的数据库迁移工具 |
| [Docker]（可选）         | 容器化平台                |

## 目录结构

\* 仅为结构，文件名未必完全一致

```shell
.
├── configs
│   ├── config.yml
│   └── config.go
├── constants
├── controllers
│   ├── auth_controller.go
│   ├── user_controller.go
│   └── ...
├── middlewares
│   ├── auth_middleware.go
│   ├── cors_middleware.go
│   └── ...
├── models
│   ├── user.go
│   └── ...
├── repositories
│   ├── auth_repository.go
│   ├── user_repository.go
│   └── ...
├── routes
│   ├── auth_routes.go
│   ├── user_routes.go
│   └── ...
├── services
│   ├── auth_service.go
│   ├── user_service.go
│   └── ...
├── utils
│   ├── response.go
│   └── ...
├── .env
├── go.mod
├── go.sum
└── main.go

```

| 目录 / 文件   | 描述                                   | 备注                                                                                 |
| ------------- | -------------------------------------- | ------------------------------------------------------------------------------------ |
| configs/      | 包含应用程序的配置文件。               | 这个目录中的文件通常用于定义应用程序的环境变量、数据库连接信息、认证密钥等等。       |
| constants/    | 包含常量文件。                         |                                                                                      |
| controllers/  | 包含处理 HTTP 请求的控制器代码。       | 这些控制器负责验证请求参数、调用适当的服务和返回响应数据。                           |
| middlewares/  | 包含处理 HTTP 请求和响应的中间件代码。 | 这些中间件可以用于实现身份验证、请求日志、跨域资源共享 (CORS) 等功能。               |
| models/       | 包含应用程序的数据模型代码。           | 这些模型代表了应用程序中的业务实体，例如用户、文章等等。                             |
| repositories/ | 包含访问数据库的代码。                 | 这些库通常包含与数据存储相关的逻辑，例如数据查询、更新、删除等等。                   |
| routes/       | 包含定义 HTTP 路由的代码。             | 这些路由定义了 API 的端点和对应的控制器和中间件。                                    |
| services/     | 包含处理应用程序逻辑的服务代码。       | 这些服务通常包含业务逻辑，例如用户验证、文章发布等等。                               |
| utils/        | 包含应用程序中的通用工具函数和结构体。 |                                                                                      |
| .env          | 包含环境变量的定义。                   | 这个文件用于定义应用程序的一些配置信息，例如数据库连接字符串、端口号、认证密钥等等。 |

[Golang]: https://go.dev/
[Gin]: https://gin-gonic.com/
[GORM]: https://gorm.io/
[PostgreSQL]: https://www.postgresql.org/docs/current/
[golang-migrate/migrate]: https://github.com/golang-migrate/migrate
[Docker]: https://www.docker.com/
