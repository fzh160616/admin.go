package router

import (
	"github.com/fzh160616/admin.go/internal/handler"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/healthz", handler.Health)

	return r
}
