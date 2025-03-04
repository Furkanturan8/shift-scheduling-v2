package dto // UpdateCragRequest Update Model
import (
	"github.com/google/uuid"
	"shift-scheduling-V2/internal/domain/entities/user"
)

// UpdateUserRequest kullanıcı güncelleme isteği için DTO
type UpdateUserRequest struct {
	ID       uuid.UUID     `json:"id,omitempty"`
	Name     string        `json:"name,omitempty"`
	Surname  string        `json:"surname,omitempty"`
	UserName string        `json:"username,omitempty"`
	Email    string        `json:"email,omitempty" validate:"omitempty,email"`
	Phone    string        `json:"phone,omitempty"`
	Password string        `json:"password,omitempty" validate:"omitempty,min=6"`
	Role     user.UserRole `json:"role,omitempty"`
}
