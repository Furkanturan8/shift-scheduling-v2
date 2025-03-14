package service

import (
	"context"
	"shift-scheduling-v2/internal/model"
	"shift-scheduling-v2/internal/repository"
	"shift-scheduling-v2/pkg/errorx"
	"time"
)

type ShiftService struct {
	shiftRepo *repository.ShiftRepository
}

func NewShiftService(shiftRepo *repository.ShiftRepository) *ShiftService {
	return &ShiftService{shiftRepo: shiftRepo}
}

func (s *ShiftService) ResetShiftsForMonth(ctx context.Context, year int, month int, locationID int) error {
	shiftStatus, err := s.shiftRepo.GetShiftStatus(ctx, year, month, locationID)
	if err != nil {
		return err
	}
	if !shiftStatus.Done {
		return errorx.ErrNotFound
	}

	shiftStatus.Done = false
	if err = s.shiftRepo.UpdateShiftStatus(ctx, shiftStatus); err != nil {
		return err
	}

	if err = s.shiftRepo.DeleteShiftsForMonth(ctx, year, month, locationID); err != nil {
		return err
	}

	return nil
}

func (s *ShiftService) IsDoctorAssignedToShift(ctx context.Context, doctorID int64, shiftDate time.Time) (bool, error) {
	return s.shiftRepo.IsDoctorAssignedToShift(ctx, doctorID, shiftDate)
}

func (s *ShiftService) GetShiftStatus(ctx context.Context, year int, month int, locationID int) (*model.ShiftsStatus, error) {
	return s.shiftRepo.GetShiftStatus(ctx, year, month, locationID)
}

func (s *ShiftService) MarkShiftStatusAsDone(ctx context.Context, year int, month int, locationID int) error {
	shiftStatus, err := s.shiftRepo.GetShiftStatus(ctx, year, month, locationID)
	if err != nil {
		return err
	}

	shiftStatus.Done = true
	return s.shiftRepo.UpdateShiftStatus(ctx, shiftStatus)
}

func (s *ShiftService) CreateShift(ctx context.Context, shift model.Shift) error {
	return s.shiftRepo.Create(ctx, shift)
}

func (s *ShiftService) GetShiftByDate(ctx context.Context, date time.Time) (*model.Shift, error) {
	return s.shiftRepo.GetShiftByDate(ctx, date)
}

func (s *ShiftService) GetTodayShifts(ctx context.Context, date time.Time) ([]model.Shift, error) {
	return s.shiftRepo.GetTodayShifts(ctx, date)
}

func (s *ShiftService) GetAllShiftsWithDetails(ctx context.Context) ([]model.Shift, error) {
	return s.shiftRepo.GetAllShiftsWithDetails(ctx)
}

func (s *ShiftService) GetShiftsByLocationID(ctx context.Context, locationID int64, month int64, year int64) ([]model.Shift, error) {
	return s.shiftRepo.GetShiftsByLocationID(ctx, locationID, month, year)
}

func (s *ShiftService) GetAllShift(ctx context.Context) (*[]model.Shift, error) {
	return s.shiftRepo.GetAllShift(ctx)
}

func (s *ShiftService) GetShiftByID(ctx context.Context, id int64) (*model.Shift, error) {
	return s.shiftRepo.GetShiftByID(ctx, id)
}

func (s *ShiftService) DeleteShift(ctx context.Context, id int64) error {
	return s.shiftRepo.DeleteShift(ctx, id)
}

func (s *ShiftService) UpdateShift(ctx context.Context, shift model.Shift) error {
	return s.shiftRepo.UpdateShift(ctx, shift)
}

func (s *ShiftService) GetShiftsStatus(ctx context.Context) ([]model.ShiftsStatus, error) {
	return s.shiftRepo.GetShiftsStatus(ctx)
}

func (s *ShiftService) GetShiftLocations(ctx context.Context) ([]model.ShiftLocation, error) {
	return s.shiftRepo.GetShiftLocations(ctx)
}

func (s *ShiftService) GetDoctorsByLocation(ctx context.Context, locationID int64) ([]model.Doctor, error) {
	return s.shiftRepo.GetDoctorsByLocation(ctx, locationID)
}

func (s *ShiftService) AssignShiftsForMonth(ctx context.Context, doctors []model.Doctor, locationID int, startOfMonth time.Time, endOfMonth time.Time) error {
	return s.shiftRepo.AssignShiftsForMonth(ctx, doctors, locationID, startOfMonth, endOfMonth)
}
