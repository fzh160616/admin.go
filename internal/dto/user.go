package dto

type UserListQuery struct {
	Keyword  string `form:"keyword"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}

type UserItem struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	TwoFAEnabled bool   `json:"two_fa_enabled"`
	LastLoginAt  any    `json:"last_login_at"`
}
