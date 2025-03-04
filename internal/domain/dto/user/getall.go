package dto

import (
	"github.com/google/uuid"
	"shift-scheduling-V2/internal/domain/entities/user"
)

// GetAllUsersResponse tüm kullanıcıları döndürmek için DTO
type GetAllUsersResponse struct {
	ID       uuid.UUID     `json:"id"`
	Name     string        `json:"name"`
	Surname  string        `json:"surname"`
	UserName string        `json:"username"`
	Email    string        `json:"email"`
	Phone    string        `json:"phone"`
	Role     user.UserRole `json:"role"`
}
