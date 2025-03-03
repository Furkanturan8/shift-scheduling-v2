package dto

import (
	"github.com/google/uuid"
	"time"
)

// GetAllUsersResult is the result of the GetAllUsersRequest Query
type GetAllUsersResult struct {
	ID        uuid.UUID
	Name      string
	Desc      string
	Country   string
	CreatedAt time.Time
}
type GetAllUserRequest struct{}
