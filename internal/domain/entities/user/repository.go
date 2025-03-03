package user

import (
	"github.com/google/uuid"
)

// Repository Interface for crags
type Repository interface {
	GetByID(id uuid.UUID) (*User, error)
	GetAll() ([]User, error)
	Add(user User) error
	Update(user User) error
	Delete(id uuid.UUID) error
}
