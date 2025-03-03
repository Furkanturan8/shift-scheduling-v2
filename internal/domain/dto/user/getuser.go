package dto

import (
	"github.com/google/uuid"
	"time"
)

// GetUserRequest Model of the Handler
type GetUserRequest struct {
	UserID uuid.UUID `json:"id" validate:"required"`
}

// GetUserResult is the return model of User Query Handlers
type GetUserResult struct {
	ID        uuid.UUID
	Name      string
	Desc      string
	Country   string
	CreatedAt time.Time
}
