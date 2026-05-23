# gin-basic

基于 [Gin](https://github.com/gin-gonic/gin) 的 Go Web 项目起步模版，提供清晰的分层结构、配置管理、日志、中间件，以及 MySQL / PostgreSQL / Redis 连接池初始化，适合作为新项目的骨架快速扩展。

## 特性

- **分层架构**：Controller → Service → Ports → Infras，职责清晰
- **依赖注入**：通过 `bootstrap.Container` 手动组装，无额外 DI 框架
- **配置管理**：Viper + YAML + `.env` + 环境变量覆盖
- **结构化日志**：Logrus，支持控制台 / 文件双输出与日志轮转
- **中间件**：Trace ID、访问日志、CORS、JWT 鉴权（可挂载）
- **基础设施**：MySQL、PostgreSQL、Redis 连接池，`host` 为空时自动跳过
- **统一响应**：`core.Result` + `core.AppError` 错误分类
- **优雅关停**：监听系统信号，释放 HTTP 与基础设施连接

## 技术栈

| 类别 | 选型 |
|------|------|
| Web 框架 | Gin |
| 配置 | Viper + godotenv |
| 日志 | Logrus + lumberjack |
| 数据库 | database/sql（MySQL / PostgreSQL） |
| 缓存 | go-redis/v9 |
| 鉴权 | golang-jwt/jwt/v5 |

## 环境要求

- Go **1.26+**
- （可选）MySQL 8+ / PostgreSQL 14+ / Redis 6+

## 快速开始

### 1. 获取代码

```bash
git clone <your-repo-url> my-app
cd my-app
```

### 2. 安装依赖

```bash
go mod download
```

### 3. 配置

复制环境变量示例文件并按需修改：

```bash
cp .env.example .env
```

主要配置文件位于 `conf/config.yml`。环境变量会覆盖 YAML 中的同名配置，规则为：配置路径中的 `.` 替换为 `_` 并转为大写。

例如：

| 配置项 | 环境变量 |
|--------|----------|
| `app.env` | `APP_ENV` |
| `database.mysql.host` | `DATABASE_MYSQL_HOST` |
| `cache.redis.host` | `CACHE_REDIS_HOST` |
| `jwt.secret` | `JWT_SECRET` |

> **说明**：数据库与 Redis 的 `host` 留空时，应用会跳过对应组件的初始化，便于本地无依赖启动。

### 4. 运行

```bash
go run ./cmd/main.go
```

默认监听 `8080` 端口（可在 `conf/config.yml` 或 `APP_API_PORT` 中修改）。

### 5. 验证

```bash
curl http://127.0.0.1:8080/api/health
```

预期响应：

```json
{
  "code": "0",
  "message": "success",
  "data": "ok"
}
```

## 项目结构

```
.
├── app/                    # 应用生命周期（启动、关停）
├── app.Dockerfile          # 应用镜像构建文件
├── cmd/
│   └── main.go             # 程序入口
├── conf/
│   └── config.yml          # 主配置文件
├── config/                 # 配置加载（Viper 封装）
├── docker-compose.yml      # 本地/演示用编排（App + MySQL + Redis）
├── internal/
│   ├── bootstrap/          # 容器组装、路由注册、基础设施初始化
│   ├── cfg/                # 配置结构体定义
│   ├── core/               # 统一响应与业务错误
│   ├── infras/             # 基础设施实现（DB、Cache）
│   │   ├── cache/
│   │   └── db/
│   ├── middleware/         # HTTP 中间件
│   ├── pkg/                # 内部公共包（trace 等）
│   ├── ports/              # 接口定义（面向依赖倒置）
│   │   ├── icache/
│   │   ├── idb/
│   │   └── iservice/
│   ├── service/            # 业务逻辑
│   ├── utils/              # 工具（JWT 等）
│   └── web/                # HTTP Controller
└── utils/
    └── logger/             # 日志初始化
```

## 配置说明

`conf/config.yml` 主要配置项：

```yaml
app:
  env: dev
  api:
    port: 8080

log:
  level: info
  output: both          # stdout | file | both
  file_path: ./logs/app.Log

database:
  connect_pool:
    max_open_conns: 100
    max_idle_conns: 20
    conn_max_lifetime: 300   # 秒
  mysql:
    host: ""             # 留空则不初始化
    database: gin_basic
  postgresql:
    host: ""

cache:
  redis:
    host: ""

jwt:
  secret: change_me
```

## 开发指南

### 新增一个 API

以「用户模块」为例，推荐按以下顺序扩展：

1. **定义接口** — `internal/ports/iservice/user.go`
2. **实现业务** — `internal/service/user.go`
3. **编写 Controller** — `internal/web/user.go`
4. **注册依赖** — 在 `internal/bootstrap/` 的 `services.go`、`controller.go` 中组装
5. **注册路由** — 在 `internal/bootstrap/router.go` 中添加

Controller 中推荐使用统一响应：

```go
func (c *UserController) Get(ctx *gin.Context) {
    data, err := c.UserService.GetByID(ctx, id)
    core.ToResult(ctx, data, err)
}
```

### 使用数据库连接池

基础设施实例挂载在 `Container` 上，Service 层通过构造函数注入：

```go
// MySQL
rows, err := c.MySQL.DB().QueryContext(ctx, "SELECT ...")

// PostgreSQL
err := c.PostgreSQL.DB().QueryRowContext(ctx, "SELECT 1").Scan(&n)

// Redis
err := c.Redis.Client().Set(ctx, "key", "value", time.Minute).Err()
```

### 接入 GORM（可选）

可复用现有 `*sql.DB` 连接池，避免重复建连：

```go
import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

gormDB, err := gorm.Open(mysql.New(mysql.Config{
    Conn: mysql.DB(), // idb.IMySQL 返回的连接池
}), &gorm.Config{})
```

建议在 `internal/infras/db/` 封装 GORM 初始化，并将 `*gorm.DB` 挂载到 `Container`。

### 启用 JWT 鉴权

项目已提供 `middleware.TokenAuth`，在路由组上挂载即可：

```go
auth := router.Group("")
auth.Use(middleware.TokenAuth(c.JwtUtils))
auth.GET("/profile", c.UserController.Profile)
```

### 中间件说明

| 中间件 | 作用 |
|--------|------|
| `Trace()` | 生成或透传 `X-Trace-Id`，写入 Context |
| `AccessLog()` | 结构化 HTTP 访问日志 |
| `Cors()` | 跨域处理（生产环境建议改为白名单） |
| `TokenAuth()` | Bearer Token JWT 鉴权 |

## 构建

```bash
# 编译
go build -o bin/gin-basic ./cmd/main.go

# 运行二进制
./bin/gin-basic
```

## Docker 部署

### 前置要求

- [Docker](https://docs.docker.com/get-docker/) 24+
- [Docker Compose](https://docs.docker.com/compose/) v2+

### 方式一：Docker Compose 一键启动（推荐）

默认启动 **App + MySQL + Redis** 完整栈：

```bash
# 构建并后台启动
docker compose up --build -d

# 查看日志
docker compose logs -f app

# 健康检查
curl http://127.0.0.1:8080/api/health

# 停止并移除容器
docker compose down

# 停止并移除容器及数据卷（清空 MySQL / Redis 数据）
docker compose down -v
```

Compose 会通过环境变量覆盖应用配置，容器内服务地址示例：

| 环境变量 | 容器内值 | 说明 |
|----------|----------|------|
| `DATABASE_MYSQL_HOST` | `mysql` | Docker 网络内的 MySQL 服务名 |
| `CACHE_REDIS_HOST` | `redis` | Docker 网络内的 Redis 服务名 |
| `LOG_OUTPUT` | `stdout` | 容器内建议输出到标准输出 |

可选：按需创建 `.env` 覆盖 Compose 变量（与应用的 `.env.example` 可共用部分键名）：

```bash
# docker compose 常用变量
APP_PORT=8080
MYSQL_ROOT_PASSWORD=root
MYSQL_DATABASE=gin_basic
JWT_SECRET=change_me_in_production
```

### 方式二：仅构建并运行应用镜像

不依赖 Compose 中的 MySQL / Redis，适合快速验证镜像或对接外部基础设施：

```bash
# 构建镜像
docker build -f app.Dockerfile -t gin-basic:latest .

# 运行（跳过 DB / Redis 初始化：host 留空）
docker run --rm -p 8080:8080 \
  -e APP_ENV=prod \
  -e LOG_OUTPUT=stdout \
  -e DATABASE_MYSQL_HOST= \
  -e CACHE_REDIS_HOST= \
  gin-basic:latest
```

对接已有 MySQL / Redis 时，传入对应环境变量即可：

```bash
docker run --rm -p 8080:8080 \
  -e DATABASE_MYSQL_HOST=192.168.1.10 \
  -e DATABASE_MYSQL_PORT=3306 \
  -e DATABASE_MYSQL_USERNAME=root \
  -e DATABASE_MYSQL_PASSWORD=secret \
  -e DATABASE_MYSQL_DATABASE=gin_basic \
  -e CACHE_REDIS_HOST=192.168.1.11 \
  -e CACHE_REDIS_PORT=6379 \
  gin-basic:latest
```

### 方式三：启用 PostgreSQL（可选 Profile）

`docker-compose.yml` 中 PostgreSQL 默认不随主栈启动，需要时使用 profile：

```bash
# 启动 App + MySQL + Redis + PostgreSQL
docker compose --profile postgres up --build -d

# 若要让 App 连接 PostgreSQL，追加环境变量（可在 docker-compose.yml 的 app.environment 中配置）
# DATABASE_POSTGRESQL_HOST=postgres
# DATABASE_POSTGRESQL_PORT=5432
# DATABASE_POSTGRESQL_USERNAME=postgres
# DATABASE_POSTGRESQL_PASSWORD=postgres
# DATABASE_POSTGRESQL_DATABASE=gin_basic
```

### 镜像说明

`app.Dockerfile` 采用多阶段构建：

1. **构建阶段**：`golang:1.26-alpine` 编译静态二进制
2. **运行阶段**：`alpine:3.21` 最小镜像，内置 `conf/` 配置，非 root 用户运行

构建上下文已通过 `.dockerignore` 排除日志、`.env`、Git 等无关文件。

### 常用 Docker 命令

```bash
# 只重建 app 服务
docker compose build app

# 进入 app 容器
docker compose exec app sh

# 查看运行状态
docker compose ps
```

## 常见问题

**Q: Docker 启动时 App 连接 MySQL 失败？**

A: 确认 `docker compose ps` 中 `mysql`、`redis` 均为 `healthy` 后再启动 app；或使用 `docker compose up --build` 让 Compose 按 `depends_on` 等待健康检查通过。

**Q: 启动时报 MySQL / Redis 连接失败？**

A: 若本地未部署对应服务，请将 `conf/config.yml` 中对应 `host` 留空，或通过环境变量置空，应用会跳过该组件初始化。

**Q: 日志文件在哪里？**

A: 默认路径为 `./logs/app.Log`，可在 `log.file_path` 中修改。日志目录已在 `.gitignore` 中忽略。

**Q: 如何替换模块名？**

A: 将 `go.mod` 中的 `module gin-basic` 改为你的模块路径，并全局替换 import 前缀。

## 后续可扩展方向

- [ ] 接入 GORM / sqlx
- [ ] 补充单元测试与集成测试
- [ ] Makefile 与 CI 流水线
- [ ] Swagger / OpenAPI 文档
- [ ] 参数校验（validator）示例

## License

WTFPL
