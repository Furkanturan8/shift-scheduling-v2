package dto

import (
	"github.com/google/uuid"
	"shift-scheduling-V2/internal/domain/entities/user"
	"time"
)

// GetUserRequest Model of the Handler
type GetUserRequest struct {
	UserID uuid.UUID `json:"id" validate:"required"`
}

// GetUserResponse kullanıcı bilgilerini döndürmek için DTO
type GetUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToGetUserResponse User entity'sini GetUserResponse'a dönüştürür
func ToGetUserResponse(u *user.User) GetUserResponse {
	return GetUserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Surname:   u.Surname,
		UserName:  u.UserName,
		Email:     u.Email,
		Phone:     u.Phone,
		Role:      u.Role.String(),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
