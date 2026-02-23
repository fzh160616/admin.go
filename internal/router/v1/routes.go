package v1

import (
	handlerv1 "github.com/fzh160616/admin.go/internal/handler/v1"
	"github.com/fzh160616/admin.go/internal/middleware"
	"github.com/fzh160616/admin.go/internal/security"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(rg *gin.RouterGroup, db *gorm.DB, jwtSecret string, rl *security.LoginRateLimiter) {
	auth := handlerv1.NewAuthHandler(db, jwtSecret, rl)
	user := handlerv1.NewUserHandler(db)

	rg.POST("/auth/register", auth.Register)
	rg.POST("/auth/login", auth.Login)

	protected := rg.Group("")
	protected.Use(middleware.JWTAuth(jwtSecret))
	protected.GET("/users", user.List)
	protected.POST("/users", user.Create)
	protected.PUT("/users/:id", user.Update)
	protected.PATCH("/users/:id/status", user.UpdateStatus)
	protected.DELETE("/users/:id", user.Delete)
}
