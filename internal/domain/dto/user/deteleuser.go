package dto

import "github.com/google/uuid"

// DeleteUserRequest kullanıcı silme isteği için DTO
type DeleteUserRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
