package persistence

import (
	"fmt"
	"github.com/google/uuid"
	"shift-scheduling-V2/internal/common/errors"
	"shift-scheduling-V2/internal/domain/entities/user"
)

type UserMemRepository struct {
	users map[string]user.User
}

func NewUserMemRepository() user.Repository {
	users := make(map[string]user.User)
	return &UserMemRepository{users}
}

// GetByID Returns the user with the provided id
func (r *UserMemRepository) GetByID(id uuid.UUID) (*user.User, error) {
	data, ok := r.users[id.String()]
	if !ok {
		return nil, errors.ErrNotFound
	}
	return &data, nil
}

// GetAll Returns all stored users
func (r *UserMemRepository) GetAll() ([]user.User, error) {
	var values []user.User
	for _, value := range r.users {
		values = append(values, value)
	}
	return values, nil
}

// Add the provided user
func (r *UserMemRepository) Add(user user.User) error {
	r.users[user.ID.String()] = user
	return nil
}

// Update the provided user
func (r *UserMemRepository) Update(user user.User) error {
	r.users[user.ID.String()] = user
	return nil
}

// Delete the user with the provided id
func (r *UserMemRepository) Delete(id uuid.UUID) error {
	_, exists := r.users[id.String()]
	if !exists {
		return fmt.Errorf("id %v not found", id.String())
	}
	delete(r.users, id.String())
	return nil
}
