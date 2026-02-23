package store

import (
	"fmt"

	"github.com/fzh160616/admin.go/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.User{}, &model.LoginLog{}); err != nil {
		return nil, fmt.Errorf("auto migrate failed: %w", err)
	}
	return db, nil
}
