package v1

import (
	"net/http"

	"github.com/fzh160616/admin.go/internal/dto"
	"github.com/gin-gonic/gin"
)

// Login 用户登录接口占位（暂未接入数据库/Redis）
func Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	// TODO: 接入 MySQL 用户校验 + Redis Token/Session 管理
	resp := dto.LoginResponse{Token: "mock-token-for-dev"}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    resp,
	})
}
