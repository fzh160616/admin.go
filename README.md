# admin.go

后台后端项目（Go + Gin）。

## 快速开始

```bash
git clone git@github.com:fzh160616/admin.go
cd admin.go
cp .env.example .env
make run
```

默认监听端口：`8080`（可通过 `APP_PORT` 覆盖）。

## 接口

- `GET /healthz` → `{ "status": "ok" }`
- `POST /api/v1/login`（占位接口）

示例请求：

```bash
curl -X POST http://127.0.0.1:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}'
```

示例响应：

```json
{
  "code": 0,
  "message": "ok",
  "data": {
    "token": "mock-token-for-dev"
  }
}
```

## 常用命令

```bash
make run
make build
make test
```
