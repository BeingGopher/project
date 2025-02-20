# Gin+GORM
搭建了一个示例，只实现了注册+登录接口，结构还是比较清晰的。
项目目录参考如下：

当然，以下是按照 Markdown 文档格式重新组织的目录结构：

markdown
# 目录结构
├── api

├── hack

├── internal

│ ├── cmd

│ ├── consts

│ ├── controller

│ ├── dao

│ ├── logic

│ ├── model

│ │ ├── do

│ │ └── entity

│ └── service

├── manifest

├── resource

├── utility

├── go.mod

└── main.go

这个目录结构展示了一个典型的 Go 项目布局，其中：
 
- `api` 目录可能包含 API 定义或接口文档。
- `hack` 目录通常用于存放一些脚本或工具，这些脚本或工具对于项目的构建、测试或部署等过程有帮助。
- `internal` 目录包含了项目的核心代码，遵循 Go 的 `internal` 包规则，这些包不能被项目外部直接导入。
  - `cmd` 目录通常存放项目的主命令或子命令的可执行文件。
  - `consts` 目录用于存放常量定义。
  - `controller` 目录包含控制器层代码，负责处理 HTTP 请求。
  - `dao` 目录包含数据访问对象（Data Access Object）层代码，负责与数据库交互。
  - `logic` 目录包含业务逻辑层代码。
  - `model` 目录包含数据模型定义，通常分为 `do`（数据传输对象）和 `entity`（实体）。
  - `service` 目录包含服务层代码，封装业务逻辑供控制器调用。
- `manifest` 目录可能包含一些配置文件或清单文件。
- `resource` 目录用于存放静态资源，如图片、模板文件等。
- `utility` 目录包含一些工具函数或帮助函数。
- `go.mod` 文件是 Go 项目的模块定义文件，用于管理项目依赖。
- `main.go` 文件是项目的入口文件。
这样的格式更适合在 Markdown 文档中展示目录结构。
