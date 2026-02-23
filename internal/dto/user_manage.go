package dto

type CreateUserRequest struct {
	Username  string `json:"username" binding:"required,min=3"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
	Enable2FA bool   `json:"enable_2fa"`
}

type UpdateUserRequest struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
}

type UpdateUserStatusRequest struct {
	Enabled bool `json:"enabled"`
}
