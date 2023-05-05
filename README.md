# 喵玉殿新版论坛服务端

开发文档：<https://root.nsdn.club/nya-develop-document/>

## 启动项目

### 下载依赖

```shell
go mod download
```

### 在 docker 中运行 MongoDB 容器（可选）

\* 如果你使用本地安装的 MongoDB 可以不需要这一步

```shell
docker compose up --detach
```

### 运行项目

\* 需要已经启动 MongoDB 了才能正常运行

```shell
go run .
```

## 目录结构

\* 仅为结构，文件名未必完全一致

```shell
.
├── configs
│   ├── config.yml
│   └── config.go
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

- `configs/`: 包含应用程序的配置文件。这个目录中的文件通常用于定义应用程序的环境变量、数据库连接信息、认证密钥等等。
- `controllers/`: 包含处理 HTTP 请求的控制器代码。这些控制器负责验证请求参数、调用适当的服务和返回响应数据。
- `middlewares/`: 包含处理 HTTP 请求和响应的中间件代码。这些中间件可以用于实现身份验证、请求日志、跨域资源共享 (CORS) 等功能。
- `models/`: 包含应用程序的数据模型代码。这些模型代表了应用程序中的业务实体，例如用户、文章等等。
- `repositories/`: 包含访问数据库的代码。这些库通常包含与数据存储相关的逻辑，例如数据查询、更新、删除等等。
- `routes/`: 包含定义 HTTP 路由的代码。这些路由定义了 API 的端点和对应的控制器和中间件。
- `services/`: 包含处理应用程序逻辑的服务代码。这些服务通常包含业务逻辑，例如用户验证、文章发布等等。
- `utils/`: 包含应用程序中的通用工具函数和结构体。
- `.env`: 包含环境变量的定义。这个文件用于定义应用程序的一些配置信息，例如数据库连接字符串、端口号、认证密钥等等。
