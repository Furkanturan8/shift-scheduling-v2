package dto

import "shift-scheduling-V2/internal/domain/entities/user"

// AddUserRequest kullanıcı oluşturma isteği için DTO
type AddUserRequest struct {
	Name     string        `json:"name" validate:"required"`
	Surname  string        `json:"surname" validate:"required"`
	UserName string        `json:"username" validate:"required"`
	Email    string        `json:"email" validate:"required,email"`
	Phone    string        `json:"phone" validate:"required"`
	Password string        `json:"password" validate:"required,min=6"`
	Role     user.UserRole `json:"role" validate:"required"`
}
