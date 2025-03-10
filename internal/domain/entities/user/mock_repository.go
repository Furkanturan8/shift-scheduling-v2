package user

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockRepository mocks Repository for testing purposes
type MockRepository struct {
	mock.Mock
}

// GetByID mock
func (m MockRepository) GetByID(id uuid.UUID) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

// GetAll mock
func (m MockRepository) GetAll() ([]User, error) {
	args := m.Called()
	return args.Get(0).([]User), args.Error(1)
}

// Add mock
func (m MockRepository) Add(user User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Update mock
func (m MockRepository) Update(user User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Delete mock
func (m MockRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
