
md_content = """
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

## API 中间件（JWT认证，日志中间件，限流中间件）
- **路径**：`router/middlewares/middlewares.go`

## mock 请求三方数据 使用全局协程池
- **路径**：`global/worker_pool.go`

## 测试接口

### 用户认证

- **POST /auth/login**：用户登录，获取 JWT Token
    - 请求体：
    ```json
    {
      "username": "example",
      "password": "password"
    }
    ```
    - 请求命令：
    ```bash
    curl --location 'http://0.0.0.0:6688/auth/login' \
      --header 'Content-Type: application/json' \
      --data '{
        "username":"example",
        "password":"password"
      }'
    ```
    - 返回：
    ```json
    {
      "token": "jwt_token"
    }
    ```

### 中间件

- **日志中间件**：自动记录请求的详细信息，帮助追踪问题。
    - 格式：
    ```text
    [2024-12-07 16:13:01] UserID: 3 - POST /tasks - 5.33975ms
    ```

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
      }'
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

- **执行任务**：
    ```bash
    curl --location --request POST 'http://0.0.0.0:6688/tasks/8/translate' \
      --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjExMjEyMzQzMiIsInVzZXJfaWQiOjMsImV4cCI6MTczMzY0MDgyNywiaXNzIjoiR29Qb2x5Z2xvdCJ9.PeJNWgi0u9gVOOjjZSeZOh-qORK5O4rgyqFBfUftSVg' \
      --data '{
        "action": "start_translation"
      }'
    ```


- **查看任务**：
    ```bash
    curl --location --request POST 'http://0.0.0.0:6688/tasks/8' \
      --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjExMjEyMzQzMiIsInVzZXJfaWQiOjMsImV4cCI6MTczMzY0MDgyNywiaXNzIjoiR29Qb2x5Z2xvdCJ9.PeJNWgi0u9gVOOjjZSeZOh-qORK5O4rgyqFBfUftSVg' \
      --data '{
        "action": "start_translation"
      }'
    ```


- **下载翻译结果**：
    ```bash
    curl --location --request POST 'http://0.0.0.0:6688/tasks/8/download' \
      --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjExMjEyMzQzMiIsInVzZXJfaWQiOjMsImV4cCI6MTczMzY0MDgyNywiaXNzIjoiR29Qb2x5Z2xvdCJ9.PeJNWgi0u9gVOOjjZSeZOh-qORK5O4rgyqFBfUftSVg' \
      --data '{
        "action": "start_translation"
      }'
    ```

