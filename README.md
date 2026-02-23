# admin.go

后台后端项目（Go + Gin）。

## 功能

- 用户注册：支持用户名 / 邮箱 / 手机号
- 用户登录：支持用户名 / 邮箱 / 手机号 + 密码
- 2FA：基于 TOTP（Google Authenticator / Microsoft Authenticator 兼容）
- 登录审计：记录每一次登录成功/失败日志（IP、UA、原因）
- 最后登录时间：登录成功后更新 `last_login_at`
- 登录限流：按 `IP + account` 维度限制连续失败次数（默认 10 分钟内 5 次失败后封禁 15 分钟）

## 快速开始

```bash
git clone git@github.com:fzh160616/admin.go
cd admin.go
cp .env.example .env
make run
```

默认监听端口：`8080`（可通过 `APP_PORT` 覆盖）。

## 接口

- `GET /healthz`
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`
- `GET /api/v1/users?page=1&page_size=10&keyword=`（用户列表/分页/搜索）

### 注册示例

```bash
curl -X POST http://127.0.0.1:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username":"admin",
    "email":"admin@example.com",
    "phone":"13800138000",
    "password":"123456",
    "enable_2fa":true
  }'
```

> 当 `enable_2fa=true` 时，响应里会带 `otp_auth_url`，用于扫码绑定认证器。

### 登录示例（用户名/邮箱/手机号都可放在 account）

```bash
curl -X POST http://127.0.0.1:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "account":"admin@example.com",
    "password":"123456",
    "two_fa_code":"123456"
  }'
```

## 安全相关环境变量

```env
LOGIN_MAX_ATTEMPTS=5
LOGIN_ATTEMPT_WINDOW_MIN=10
LOGIN_BLOCK_MIN=15
```

## 常用命令

```bash
make run
make build
make test
```
