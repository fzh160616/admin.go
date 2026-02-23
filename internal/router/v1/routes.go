package v1

import (
	handlerv1 "github.com/fzh160616/admin.go/internal/handler/v1"
	"github.com/fzh160616/admin.go/internal/security"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(rg *gin.RouterGroup, db *gorm.DB, jwtSecret string, rl *security.LoginRateLimiter) {
	auth := handlerv1.NewAuthHandler(db, jwtSecret, rl)
	rg.POST("/auth/register", auth.Register)
	rg.POST("/auth/login", auth.Login)
}
