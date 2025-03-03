package user

import (
	"github.com/google/uuid"
	"time"
)

// Crag Model that represents the Crag
type User struct {
	ID        uuid.UUID
	Name      string
	Desc      string
	Country   string
	CreatedAt time.Time
}
