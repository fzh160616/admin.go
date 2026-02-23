package dto

type RegisterRequest struct {
	Username  string `json:"username" binding:"required,min=3"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
	Enable2FA bool   `json:"enable_2fa"`
}

type LoginRequest struct {
	Account   string `json:"account" binding:"required"` // username/email/phone
	Password  string `json:"password" binding:"required"`
	TwoFACode string `json:"two_fa_code"`
}
