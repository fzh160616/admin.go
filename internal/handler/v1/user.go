package v1

import (
	"net/http"
	"strings"

	"github.com/fzh160616/admin.go/internal/dto"
	"github.com/fzh160616/admin.go/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) List(c *gin.Context) {
	var q dto.UserListQuery
	_ = c.ShouldBindQuery(&q)
	if q.Page <= 0 {
		q.Page = 1
	}
	if q.PageSize <= 0 || q.PageSize > 100 {
		q.PageSize = 10
	}

	dbq := h.db.Model(&model.User{})
	if kw := strings.TrimSpace(q.Keyword); kw != "" {
		like := "%" + kw + "%"
		dbq = dbq.Where("username LIKE ? OR email LIKE ? OR phone LIKE ?", like, like, like)
	}

	var total int64
	if err := dbq.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "query total failed"})
		return
	}

	var users []model.User
	if err := dbq.Order("id desc").Offset((q.Page - 1) * q.PageSize).Limit(q.PageSize).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "query users failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data": gin.H{
			"list":      users,
			"total":     total,
			"page":      q.Page,
			"page_size": q.PageSize,
		},
	})
}
