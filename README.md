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

## 常用命令

```bash
make run
make build
make test
```
