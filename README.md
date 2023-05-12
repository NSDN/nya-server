# 喵玉殿新版论坛服务端

开发文档：<https://root.nsdn.club/nya-develop-document/>

## 技术栈

|技术栈|功能|
|-|-|
|Docker（可选）|容器化平台|
|Golang|编程语言|
|Gin|RESTful 框架|
|MongoDB|数据库|

## 设置环境变量文件（重要！）

- 希望使用默认环境变量请将示例文件 `.env.*.example` 中的 `example` 去掉
- 希望自己生成环境变量文件请参考示例文件
- 根据下文不同的启动方式需要设置不同的环境变量

## 启动项目（本地有 Golang 环境）

### 下载依赖

```shell
# 下载依赖
go mod download
# 验证依赖
go mod verify
```

### 在 docker 中运行 MongoDB 容器（可选）

- 如果你使用本地安装的 MongoDB 可以不需要这一步
- 如果你使用本地安装的 MongoDB 记得修改 `.env` 文件中的 `MONGODB_URI`

```shell
docker compose up --detach
```

## 启动项目（本地无 Golang 环境）

- 需要本地环境中有 docker
- 这个命令会同时在 docker 中创建并运行 Golang 项目容器和 MongoDB 容器
- 需要更改 `.env` 中的 `MONGODB_URI`

```shell
# 启动项目
docker compose --file docker-comopse.no-local.go.yml up --detach

# 关闭项目
docker compose --file docker-compose.no-local-go.yml down
```

### 查看 docker 中的 Golang 项目目录

```shell
# 进入容器
# 这个命令会在当前命令行打开一个 bash 进入容器
docker compose exec forum-backend bash

# 在容器内执行 linux 命令
ls -la 
```

### 查看 docker 中的 Golagn 项目的运行日志

```shell
docker compose logs forum-backend
```

### 运行项目

\* 需要已经启动 MongoDB 了才能正常运行

```shell
go run .
```

## 验证项目已正常启动

```shell
# 返回 pong 则表示启动成功
curl -X GET 'http://localhost:10127/ping'
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

|目录 / 文件|描述|备注|
|-|-|-|
|configs/|包含应用程序的配置文件。|这个目录中的文件通常用于定义应用程序的环境变量、数据库连接信息、认证密钥等等。|
|controllers/|包含处理 HTTP 请求的控制器代码。|这些控制器负责验证请求参数、调用适当的服务和返回响应数据。|
|middlewares/|包含处理 HTTP 请求和响应的中间件代码。|这些中间件可以用于实现身份验证、请求日志、跨域资源共享 (CORS) 等功能。|
|models/|包含应用程序的数据模型代码。|这些模型代表了应用程序中的业务实体，例如用户、文章等等。|
|repositories/|包含访问数据库的代码。|这些库通常包含与数据存储相关的逻辑，例如数据查询、更新、删除等等。|
|routes/|包含定义 HTTP 路由的代码。|这些路由定义了 API 的端点和对应的控制器和中间件。|
|services/|包含处理应用程序逻辑的服务代码。|这些服务通常包含业务逻辑，例如用户验证、文章发布等等。|
|utils/|包含应用程序中的通用工具函数和结构体。||
|.env|包含环境变量的定义。|这个文件用于定义应用程序的一些配置信息，例如数据库连接字符串、端口号、认证密钥等等。|
