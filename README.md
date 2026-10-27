# 光影捕梦阁 · 摄影博客系统

一个用 Go 语言实现的摄影主题博客系统：支持用户注册登录、文章发布与管理、评论互动、文章搜索，页面以服务端模板渲染，整体走「MVC + 分层」结构。

> 项目讲解视频未随仓库上传（体积约 294 MB，超过 GitHub 单文件 100 MB 的限制），如需查看请另行索取或观看分享链接。

---

## 功能

| 模块 | 说明 |
| --- | --- |
| 首页 | 文章列表展示（`/`） |
| 文章 | 查看详情 `/articles/:id`、按标题搜索 `/articles/search` |
| 文章管理 | 写文章、编辑、删除（`/articles/create`、`/:id/edit`、`/:id/update`、`DELETE /:id`），**需登录** |
| 用户 | 注册 `/register`、登录 `/login`、退出 `/logout`，**游客可访问** |
| 个人中心 | 查看与修改个人信息（`/information`、`/information/edit`），**需登录** |
| 评论 | 发表评论 `POST /comments`，**需登录** |

权限通过两个中间件控制：`middlewares.IsAuth()`（已登录才可访问）、`middlewares.IsGuest()`（未登录才可访问，防止重复登录）。

---

## 技术栈

- **语言**：Go 1.18
- **Web 框架**：[Gin](https://gin-gonic.com/) 1.9
- **ORM / 数据库**：[GORM](https://gorm.io/) 1.24 + MySQL 驱动
- **配置**：[Viper](https://github.com/spf13/viper)（读取 `config/config.yml`）
- **命令行**：[Cobra](https://github.com/spf13/cobra)（`serve` / `migrate` / `seed` 三个子命令）
- **会话**：gin-contrib/sessions（基于 gorilla/sessions + securecookie）
- **密码加密**：bcrypt（`golang.org/x/crypto`）
- **页面**：Go 标准库 `html/template` + Bootstrap 静态资源

---

## 目录结构

```
.
├── lmt_blog.sql                 # 数据库结构与初始数据（users / articles / comments）
└── lmt_BlogProject/blog/        # 项目根目录（module 名为 blog）
    ├── main.go                  # 入口，调用 cmd.Execute()
    ├── go.mod / go.sum
    ├── cmd/                     # Cobra 子命令：serve / migrate / seed
    ├── config/
    │   ├── config.go            # 配置结构体
    │   └── config.yml           # 实际配置（服务地址、数据库账号）
    ├── internal/                # 业务代码
    │   ├── database/            # migration（建表）、seeder（演示数据）
    │   ├── middlewares/         # IsAuth / IsGuest 鉴权中间件
    │   ├── modules/             # article / comment / home / user 四个模块
    │   │   └── <模块>/{controllers,services,repositories,models,requests,responses,routes,html}
    │   ├── providers/           # 路由注册、校验错误、模板渲染
    │   └── templates/           # layouts（head/header/navbar/footer）与错误页
    ├── pkg/                     # 可复用基础设施
    │   ├── bootstrap/           # serve / migrate / seed 的启动流程
    │   ├── config/              # 配置读写封装
    │   ├── database/            # GORM 连接
    │   ├── routing/             # Gin 引擎初始化与启动
    │   ├── sessions/ static/ html/
    │   └── converters/ errors/
    └── assets/                  # 静态资源：css / js / img / demopic（示例配图）
```

模块内部按 **controller → service → repository → model** 分层，service 与 repository 均定义了接口，便于替换实现。

---

## 本地运行

### 环境要求

- Go >= 1.18
- MySQL >= 5.7（需支持 `utf8mb4`）

### 步骤

**1. 创建数据库并导入初始数据**

```bash
# 登录 MySQL 后执行
CREATE DATABASE blog DEFAULT CHARACTER SET utf8mb4;

# 退出 MySQL，在仓库根目录执行（按提示输入密码）
mysql -u root -p blog < lmt_blog.sql
```

也可以不导入 SQL，直接用程序建表（二选一即可）：

```bash
cd lmt_BlogProject/blog
go run main.go migrate    # GORM AutoMigrate 建 users / articles / comments 三张表
```

**2. 修改数据库配置**

编辑 `lmt_BlogProject/blog/config/config.yml`，改成你自己的 MySQL 账号密码：

```yaml
server:
  host: "localhost"
  port: "8080"

db:
  username: "root"
  password: "你的密码"
  host: "127.0.0.1"
  port: "3306"
  name: "blog"
```

> 注意：Viper 是按**相对于当前工作目录**的 `config` 路径读取配置的，所以下面的命令必须**先 `cd` 到 `blog` 目录**再执行，否则会报 `Error reading the configs`。

**3. 启动服务**

```bash
cd lmt_BlogProject/blog
go mod download
go run main.go serve
```

浏览器打开 <http://localhost:8080> 即可。

**4.（可选）生成演示数据**

```bash
go run main.go seed
```

会创建一个演示账号 `random@email.com`（密码 `secret`，bcrypt 加密）并写入 10 篇示例文章。若导入 `lmt_blog.sql` 后无法登录，可删除 `users` 表中该账号再执行本命令重新生成。

### 命令一览

| 命令 | 作用 |
| --- | --- |
| `go run main.go serve` | 启动 Web 服务（读取 config.yml 中的 host/port） |
| `go run main.go migrate` | 根据 model 自动建表 |
| `go run main.go seed` | 写入演示用户与文章 |

---

## 说明

- **示例配图**：`assets/demopic/` 下的图片是演示用素材，可自行替换。
- **配置中的密码**：`config/config.yml` 里保留的是本地开发环境的默认密码，部署到任何公开环境前请务必修改，或将该文件加入 `.gitignore` 并另附 `config.yml.example`。
- **数据库脚本**：`lmt_blog.sql` 同时包含建表语句与开发过程中的真实数据，仅用于快速还原演示环境。
