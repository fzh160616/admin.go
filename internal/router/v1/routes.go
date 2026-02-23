package v1

import (
	handlerv1 "github.com/fzh160616/admin.go/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

func Register(rg *gin.RouterGroup) {
	rg.POST("/login", handlerv1.Login)
}
