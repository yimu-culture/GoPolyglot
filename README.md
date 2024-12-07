


## 安装与运行

1. 克隆该项目：
    ```bash
    git clone https://github.com/yimu-culture/GoPolyglot
    cd GoPolyglot
    ```

2. 安装 Go 依赖：
    ```bash
    go mod tidy
    ```

3. 启动应用：
    ```bash
    go run main.go
    ```

## API 中间件(jwt认证，日志中间价，限流中间件)
- `router/middlewares/middlewares.go`

## mock 请求三方数据 使用全局协程池
- `global/worker_pool.go`



## 测试接口

#### 用户认证

- **POST /auth/login**：用户登录，获取 JWT Token
    - 请求体：`{"username": "example", "password": "password"}`
    - 请求体：`curl --location 'http://0.0.0.0:6688/auth/login' \
      --header 'Content-Type: application/json' \
      --data '{
      "username":"example",
      "password":"password"
      }'`
    - 返回：`{"token": "jwt_token"}`

#### 任务管理

- **POST /tasks**：创建翻译任务
    - 请求体：`{"file": "file_content"}`
    - 返回：任务 ID 和状态。

- **POST /tasks/{task_id}/translate**：执行翻译任务
    - 请求头：`Authorization: Bearer {jwt_token}`
    - 返回：任务开始的状态信息。

- **GET /tasks/{task_id}**：查询任务状态
    - 返回：任务的当前状态、进度和翻译结果。

- **GET /tasks/{task_id}/download**：下载翻译文件
    - 返回：翻译后的文件内容。

### 中间件

- **日志中间件**：自动记录请求的详细信息，帮助追踪问题。
    - 格式：`[2024-11-21 12:00:00] UserID: 123e4567 - POST /tasks/create - 150ms`

- **限流中间件**：每个用户每分钟最多请求 10 次，超过限制返回 `429 Too Many Requests` 错误。

## 测试

### 示例 API 请求

- **用户注册**：
    ```bash
    curl --location 'http://0.0.0.0:6688/auth/users' \
      --header 'Content-Type: application/json' \
      --data '{
      "username":"example",
      "password":"password"
      }''
    ```

- **用户登录**：
    ```bash
    curl --location 'http://0.0.0.0:6688/auth/login' \
      --header 'Content-Type: application/json' \
      --data '{
      "username":"example",
      "password":"password"
      }'
    ```

- **创建翻译任务**：
    ```bash
  curl --location 'http://0.0.0.0:6688/tasks' \
    --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjExMjEyMzQzMiIsInVzZXJfaWQiOjMsImV4cCI6MTczMzYzODY4MiwiaXNzIjoiR29Qb2x5Z2xvdCJ9.qyWF7vHLchfjvMhl1kvJu_IvLwzL6BzIE7GNHI_Splc' \
    --header 'Content-Type: application/json' \
    --data '{
    "source_lang":"en",
    "target_lang":"ch",
    "source_doc":"r34t24"
    }'    
  ```
