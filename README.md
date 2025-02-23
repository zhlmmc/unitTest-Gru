# Order System

这是一个基于 Go 语言开发的订单系统，采用清晰的目录结构和模块化设计。

## 项目结构

```
.
├── cmd/                    # 命令行工具和应用程序入口
│   └── server/            # 服务器应用程序入口
├── pkg/                    # 项目核心包
│   ├── platform/          # 平台相关的代码
│   └── infra/             # 基础设施代码
├── go.mod                 # Go 模块依赖管理文件
└── go.sum                 # Go 模块依赖版本锁定文件
```

## 技术栈

- Go 1.21
- MySQL (go-sql-driver/mysql v1.7.1)
- 测试框架：
  - testify v1.8.4
  - go-sqlmock v1.5.2

## 开始使用

### 前置条件

- Go 1.21 或更高版本
- MySQL 数据库

### 安装依赖

```bash
go mod download
```

### 运行服务

```bash
go run cmd/server/main.go
```

## 测试

运行所有测试：

```bash
go test ./...
```

## 项目特点

- 清晰的目录结构，遵循 Go 项目最佳实践
- 完整的单元测试支持
- 模块化设计，便于扩展和维护

## 许可证

[添加许可证信息]
