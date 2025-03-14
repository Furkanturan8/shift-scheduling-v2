package service

import (
	"context"
	"shift-scheduling-v2/internal/dto"
	"shift-scheduling-v2/internal/model"
	"shift-scheduling-v2/internal/repository"
	"shift-scheduling-v2/pkg/errorx"
)

type ShiftService struct {
	shiftRepo *repository.ShiftRepository
}

func NewShiftService(shiftRepo *repository.ShiftRepository) *ShiftService {
	return &ShiftService{shiftRepo: shiftRepo}
}

func (s *ShiftService) CreateShift(ctx context.Context, vm dto.ShiftCreateRequest) error {
	shift := vm.ToDBModel(model.Shift{})
	err := s.shiftRepo.Create(ctx, shift)
	if err != nil {
		return errorx.ErrDatabaseOperation
	}
	return nil
}
