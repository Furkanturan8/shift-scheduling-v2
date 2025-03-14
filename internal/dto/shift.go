package dto

import (
	"shift-scheduling-v2/internal/model"
	"time"
)

type ShiftCreateRequest struct {
	DoctorID   int64     `json:"doctor_id" validate:"required"`
	LocationID int64     `json:"location_id" validate:"required"`
	ShiftDate  time.Time `json:"shift_date" validate:"required"`
	StartTime  string    `json:"start_time" validate:"required"`
	EndTime    string    `json:"end_time" validate:"required"`
}

func (vm ShiftCreateRequest) ToDBModel(m model.Shift) model.Shift {
	m.DoctorID = vm.DoctorID
	m.LocationID = vm.LocationID
	m.ShiftDate = vm.ShiftDate
	m.StartTime = vm.StartTime
	m.EndTime = vm.EndTime

	return m
}

type ShiftUpdateRequest struct {
	ShiftID    int64     `json:"shift_id" validate:"required"`
	DoctorID   int64     `json:"doctor_id" validate:"required"`
	LocationID int64     `json:"location_id" validate:"required"`
	ShiftDate  time.Time `json:"shift_date" validate:"required"`
	StartTime  string    `json:"start_time" validate:"required"`
	EndTime    string    `json:"end_time" validate:"required"`
}

func (vm ShiftUpdateRequest) ToDBModel(m model.Shift) model.Shift {
	m.ID = vm.ShiftID
	m.DoctorID = vm.DoctorID
	m.LocationID = vm.LocationID
	m.ShiftDate = vm.ShiftDate
	m.StartTime = vm.StartTime
	m.EndTime = vm.EndTime

	return m
}

type ShiftResponse struct {
	ID         int       `json:"id"`
	DoctorID   int64     `json:"doctor_id"`
	LocationID int64     `json:"location_id"`
	ShiftDate  time.Time `json:"shift_date"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}

func (vm ShiftResponse) ToResponseModel(m model.Shift) ShiftResponse {
	vm.ID = int(m.ID)
	vm.DoctorID = m.DoctorID
	vm.LocationID = m.LocationID
	vm.ShiftDate = m.ShiftDate
	vm.StartTime = m.StartTime
	vm.EndTime = m.EndTime

	return vm
}

type ShiftListResponse struct {
	Shifts []ShiftResponse `json:"shifts"`
	Total  int             `json:"total"`
}

func (vm ShiftListResponse) ToResponseModel(m []model.Shift) ShiftListResponse {
	shifts := make([]ShiftResponse, len(m))
	for i, shift := range m {
		shifts[i] = ShiftResponse{
			ID:         int(shift.ID),
			DoctorID:   shift.DoctorID,
			LocationID: shift.LocationID,
			ShiftDate:  shift.ShiftDate,
			StartTime:  shift.StartTime,
			EndTime:    shift.EndTime,
		}
	}

	return ShiftListResponse{
		Shifts: shifts,
		Total:  len(shifts),
	}
}
