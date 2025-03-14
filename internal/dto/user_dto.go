package dto

import (
	"shift-scheduling-v2/internal/model"
)

// rolu create ederken set ediyorum gerek yok
type UserCreateDTO struct {
	Email    string     `json:"email" validate:"required_without=Phone,omitempty,max=64,email"`
	Phone    string     `json:"phone" validate:"required_without=Email,omitempty,max=11,numeric"`
	Name     string     `json:"name" validate:"required,max=100"`
	Surname  string     `json:"surname" validate:"required,max=100"`
	Password string     `json:"password" validate:"required,min=3,max=100"`
	Role     model.Role `json:"role" validate:"required"`
}

func (vm UserCreateDTO) ToDBModel(m model.User) model.User {
	m.Email = vm.Email
	m.Phone = vm.Phone
	m.Name = vm.Name
	m.Surname = vm.Surname
	_ = m.SetPassword(vm.Password)
	m.Role = vm.Role

	return m
}

type UserResponseDTO struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Status   string `json:"active"`
}

func (vm UserResponseDTO) ToResponseModel(m model.User) UserResponseDTO {
	vm.ID = m.ID
	vm.Email = m.Email
	vm.Name = m.Name
	vm.Surname = m.Surname
	vm.Username = m.Username
	vm.Phone = m.Phone
	vm.Role = m.Role.String()
	vm.Status = string(m.Status)

	return vm
}
