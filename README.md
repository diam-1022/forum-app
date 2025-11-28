# Simple Forum - 简单论坛

一个基于 Vue 3 + Go + MySQL 的全栈论坛应用演示项目。

## 技术栈

- **前端**: Vue 3 + Vite + Pinia + Axios
  - 单页应用（SPA）
  - 支持板块/话题筛选
  - 帖子发布和评论功能
  
- **后端**: Go + Gin + GORM
  - RESTful API 接口
  - 数据持久化
  
- **数据库**: MySQL
  - 数据库结构定义在 `backend/migrations` 目录
  - 包含初始数据和种子数据

## 快速开始

### 前置要求

- Go 1.21 或更高版本
- Node.js 16 或更高版本
- MySQL 5.7 或更高版本

### 1. 初始化数据库

```bash
# 登录 MySQL 并创建数据库
mysql -u root -p

# 在 MySQL 中执行：
CREATE DATABASE forum_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
exit

# 执行数据库迁移脚本
mysql -u root -p forum_db < backend/migrations/001_init.sql
mysql -u root -p forum_db < backend/migrations/002_seed.sql
```

### 2. 配置后端环境变量

```bash
# 进入后端目录
cd backend

# 创建 .env 文件（如果不存在）
# Windows PowerShell:
@"
APP_PORT=8080
MYSQL_HOST=127.0.0.1
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=你的MySQL密码
MYSQL_DB=forum_db
MYSQL_PARAMS=charset=utf8mb4&parseTime=True&loc=Local
"@ | Out-File -FilePath .env -Encoding utf8

# Linux/Mac:
cat > .env << EOF
APP_PORT=8080
MYSQL_HOST=127.0.0.1
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=你的MySQL密码
MYSQL_DB=forum_db
MYSQL_PARAMS=charset=utf8mb4&parseTime=True&loc=Local
EOF
```

**重要**: 请将 `你的MySQL密码` 替换为你的实际 MySQL root 密码。

### 3. 启动后端服务

```bash
# 进入后端目录
cd backend

# 安装依赖
go mod tidy

# 启动服务器（默认端口 8080）
go run ./cmd/api
```

后端服务启动后，你应该看到类似以下输出：

```
[GIN-debug] Listening
```

### 4. 启动前端服务

```bash
# 进入前端目录
cd frontend

# 安装依赖
npm install

# 启动前端开发服务器
npm run dev
```

前端服务启动后，你应该看到类似以下输出：

```
[Vite] Listening on http://localhost:5173
```

### 5. 访问应用

打开浏览器，访问 http://localhost:5173

## 项目结构

```
backend/
  cmd/api/main.go          Gin entrypoint
  internal/config          env helpers
  internal/db              GORM connection + migrations
  internal/handlers        REST handlers
  internal/repository      Data access helpers
  migrations/*.sql         schema + seed
frontend/
  src/api/http.js          Axios instance
  src/stores/forumStore.js Pinia store wrapping the API
  src/components/*         Reusable UI widgets
  src/App.vue              Layout + orchestration
```



