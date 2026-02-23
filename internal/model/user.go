package model

import "time"

type User struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Username     string     `gorm:"size:64;uniqueIndex;not null" json:"username"`
	Email        string     `gorm:"size:128;uniqueIndex;not null" json:"email"`
	Phone        string     `gorm:"size:32;uniqueIndex;not null" json:"phone"`
	PasswordHash string     `gorm:"size:255;not null" json:"-"`
	TwoFAEnabled bool       `gorm:"not null;default:false" json:"two_fa_enabled"`
	TwoFASecret  string     `gorm:"size:128" json:"-"`
	Enabled      bool       `gorm:"not null;default:true" json:"enabled"`
	LastLoginAt  *time.Time `json:"last_login_at"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
