package v1

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fzh160616/admin.go/internal/dto"
	"github.com/fzh160616/admin.go/internal/model"
	"github.com/fzh160616/admin.go/internal/security"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db          *gorm.DB
	jwtSecret   string
	rateLimiter *security.LoginRateLimiter
}

func NewAuthHandler(db *gorm.DB, jwtSecret string, rl *security.LoginRateLimiter) *AuthHandler {
	return &AuthHandler{db: db, jwtSecret: jwtSecret, rateLimiter: rl}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "invalid request", "error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "password hash failed"})
		return
	}

	user := model.User{
		Username:     strings.TrimSpace(req.Username),
		Email:        strings.ToLower(strings.TrimSpace(req.Email)),
		Phone:        strings.TrimSpace(req.Phone),
		PasswordHash: string(hash),
		TwoFAEnabled: req.Enable2FA,
	}

	var otpURL string
	if req.Enable2FA {
		key, err := totp.Generate(totp.GenerateOpts{Issuer: "admin.go", AccountName: user.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "2fa setup failed"})
			return
		}
		user.TwoFASecret = key.Secret()
		otpURL = key.URL()
	}

	if err := h.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "user create failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data": gin.H{
			"id":          user.ID,
			"username":    user.Username,
			"email":       user.Email,
			"phone":       user.Phone,
			"enable_2fa":  user.TwoFAEnabled,
			"otp_auth_url": otpURL,
		},
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "invalid request", "error": err.Error()})
		return
	}

	account := strings.TrimSpace(req.Account)
	ip := c.ClientIP()
	ua := c.Request.UserAgent()

	if ok, remain := h.rateLimiter.Allow(ip, account); !ok {
		h.writeLog(nil, account, ip, ua, false, "rate limited")
		c.JSON(http.StatusTooManyRequests, gin.H{
			"code":    429,
			"message": fmt.Sprintf("too many attempts, retry after %d seconds", int(remain.Seconds())),
		})
		return
	}

	var user model.User
	err := h.db.Where("username = ? OR email = ? OR phone = ?", account, strings.ToLower(account), account).First(&user).Error
	if err != nil {
		h.rateLimiter.RecordFailure(ip, account)
		h.writeLog(nil, account, ip, ua, false, "user not found")
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		h.rateLimiter.RecordFailure(ip, account)
		h.writeLog(&user.ID, account, ip, ua, false, "wrong password")
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "invalid credentials"})
		return
	}

	if user.TwoFAEnabled {
		if req.TwoFACode == "" {
			h.rateLimiter.RecordFailure(ip, account)
			h.writeLog(&user.ID, account, ip, ua, false, "2fa required")
			c.JSON(http.StatusUnauthorized, gin.H{"code": 4011, "message": "2fa required"})
			return
		}
		if ok := totp.Validate(strings.TrimSpace(req.TwoFACode), user.TwoFASecret); !ok {
			h.rateLimiter.RecordFailure(ip, account)
			h.writeLog(&user.ID, account, ip, ua, false, "invalid 2fa code")
			c.JSON(http.StatusUnauthorized, gin.H{"code": 4012, "message": "invalid 2fa code"})
			return
		}
	}

	now := time.Now()
	if err := h.db.Model(&user).Update("last_login_at", now).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "update last login failed"})
		return
	}

	token, err := h.signJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "token sign failed"})
		return
	}

	h.rateLimiter.RecordSuccess(ip, account)
	h.writeLog(&user.ID, account, ip, ua, true, "login success")
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "ok", "data": gin.H{"token": token, "last_login_at": now}})
}

func (h *AuthHandler) signJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(h.jwtSecret))
}

func (h *AuthHandler) writeLog(userID *uint, account, ip, userAgent string, success bool, reason string) {
	_ = h.db.Create(&model.LoginLog{
		UserID:    userID,
		Account:   account,
		IP:        ip,
		UserAgent: userAgent,
		Success:   success,
		Reason:    reason,
	}).Error
}
