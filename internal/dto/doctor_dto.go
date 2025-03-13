package dto

import (
	"shift-scheduling-v2/internal/model"
	"time"
)

// create and update
type CreateDoctorDTO struct {
	UserID         int64  `json:"user_id" validate:"required"`
	Title          string `json:"title" validate:"required"`
	Specialization string `json:"specialization" validate:"required"`
}

func (vm CreateDoctorDTO) ToDBModel(m model.Doctor) model.Doctor {
	m.UserID = vm.UserID
	m.Specialization = vm.Specialization
	m.Title = vm.Title

	return m
}

type DoctorResponseDTO struct {
	ID             int64           `json:"id"`
	UserID         int64           `json:"user_id"`
	LocationID     int64           `json:"location_id"`
	Title          string          `json:"title"`
	Specialization string          `json:"specialization"`
	Status         string          `json:"status"`
	User           UserResponseDTO `json:"user,omitempty"`
}

func (vm DoctorResponseDTO) ToResponseModel(m model.Doctor) *DoctorResponseDTO {
	vm.ID = m.ID
	vm.UserID = m.UserID
	vm.Specialization = m.Specialization
	vm.Title = m.Title
	vm.User = UserResponseDTO{}.ToResponseModel(m.User)

	return &vm
}

type DoctorHolidayDTO struct {
	ID            int64     `json:"id"`
	DoctorID      int64     `json:"doctor_id"`
	LocationID    int64     `json:"location_id"`
	HolidayDate   time.Time `json:"holiday_date"`
	LocationName  string    `json:"location"`
	DoctorName    string    `json:"doctor_name"`
	DoctorSurname string    `json:"doctor_surname"`
}

func (vm DoctorHolidayDTO) ToResponseModel(m model.Holiday) DoctorHolidayDTO {
	vm.ID = m.ID
	vm.DoctorID = m.DoctorID
	vm.LocationID = m.LocationID
	vm.HolidayDate = m.HolidayDate
	vm.LocationName = m.Location.Name
	vm.DoctorName = m.Doctor.User.Name
	vm.DoctorSurname = m.Doctor.User.Surname

	return vm
}
