package dto

import "github.com/google/uuid"

// DeleteUserRequest Command Model
type DeleteUserRequest struct {
	UserID uuid.UUID `json:"id" validate:"required"`
}
