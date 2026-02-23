package router

import (
	"github.com/fzh160616/admin.go/internal/handler"
	routerv1 "github.com/fzh160616/admin.go/internal/router/v1"
	"github.com/fzh160616/admin.go/internal/security"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(db *gorm.DB, jwtSecret string, rl *security.LoginRateLimiter) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/healthz", handler.Health)

	apiV1 := r.Group("/api/v1")
	routerv1.Register(apiV1, db, jwtSecret, rl)

	return r
}
