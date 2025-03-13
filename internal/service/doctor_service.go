package service

import (
	"context"
	"shift-scheduling-v2/internal/dto"
	"shift-scheduling-v2/internal/model"
	"shift-scheduling-v2/internal/repository"
	"shift-scheduling-v2/pkg/errorx"
)

type DoctorService struct {
	doctorRepo *repository.DoctorRepository
	userRepo   *repository.UserRepository
}

func NewDoctorService(doctorRepo *repository.DoctorRepository, userRepo *repository.UserRepository) *DoctorService {
	return &DoctorService{
		doctorRepo: doctorRepo,
		userRepo:   userRepo,
	}
}

func (s *DoctorService) GetDoctorByShiftID(ctx context.Context, shiftID int64) (*dto.DoctorResponseDTO, error) {
	doctor, err := s.doctorRepo.GetByShiftID(ctx, shiftID)
	if err != nil {
		return nil, errorx.ErrNotFound
	}

	return dto.DoctorResponseDTO{}.ToResponseModel(*doctor), nil
}

func (s *DoctorService) GetDoctorsByLocation(ctx context.Context, locationID int64) ([]dto.DoctorResponseDTO, error) {
	doctors, err := s.doctorRepo.GetByLocation(ctx, locationID)
	if err != nil {
		return nil, errorx.ErrDatabaseOperation
	}

	var doctorList []dto.DoctorResponseDTO
	for _, doctor := range doctors {
		drDto := dto.DoctorResponseDTO{}.ToResponseModel(doctor)
		doctorList = append(doctorList, *drDto)
	}

	return doctorList, nil
}

func (s *DoctorService) GetDoctorHolidays(ctx context.Context, doctorID int64) ([]dto.DoctorHolidayDTO, error) {
	holidays, err := s.doctorRepo.GetHolidaysByDoctor(ctx, doctorID)
	if err != nil {
		return nil, errorx.ErrDatabaseOperation
	}

	var holidayList []dto.DoctorHolidayDTO
	for _, holiday := range holidays {
		hDto := dto.DoctorHolidayDTO{}.ToResponseModel(holiday)
		holidayList = append(holidayList, hDto)
	}

	return holidayList, nil
}

func (s *DoctorService) GetDoctorsHolidayByLocationId(ctx context.Context, locationID, month, year int64) ([]dto.DoctorHolidayDTO, error) {
	holidays, err := s.doctorRepo.GetHolidaysByLocation(ctx, locationID, month, year)
	if err != nil {
		return nil, errorx.ErrDatabaseOperation
	}

	var holidayList []dto.DoctorHolidayDTO
	for _, holiday := range holidays {
		hDto := dto.DoctorHolidayDTO{}.ToResponseModel(holiday)
		holidayList = append(holidayList, hDto)
	}

	return holidayList, nil
}

func (s *DoctorService) Create(ctx context.Context, req *dto.CreateDoctorDTO) error {
	// Kullanıcı kontrolü
	user, err := s.userRepo.GetByID(ctx, req.UserID)
	if err != nil || user.Role != model.UserRoleDoctor {
		return errorx.ErrInvalidRequest
	}

	doctor := req.ToDBModel(model.Doctor{})

	if err = s.doctorRepo.Create(ctx, &doctor); err != nil {
		return errorx.ErrDatabaseOperation
	}

	return nil
}

func (s *DoctorService) List(ctx context.Context) ([]dto.DoctorResponseDTO, error) {
	doctors, _, err := s.doctorRepo.List(ctx, "User") // todo: total count
	if err != nil {
		return nil, errorx.ErrDatabaseOperation
	}

	var doctorList []dto.DoctorResponseDTO
	for _, doctor := range doctors {
		drDto := dto.DoctorResponseDTO{}.ToResponseModel(doctor)
		doctorList = append(doctorList, *drDto)
	}

	return doctorList, nil
}

func (s *DoctorService) GetByID(ctx context.Context, id int64) (*dto.DoctorResponseDTO, error) {
	doctor, err := s.doctorRepo.GetByID(ctx, id, "User")
	if err != nil {
		return nil, errorx.ErrNotFound
	}

	return dto.DoctorResponseDTO{}.ToResponseModel(*doctor), nil
}

func (s *DoctorService) Update(ctx context.Context, id int64, req *dto.CreateDoctorDTO) error {
	doctor, err := s.doctorRepo.GetByID(ctx, id)
	if err != nil {
		return errorx.ErrNotFound
	}

	req.ToDBModel(*doctor)

	if err = s.doctorRepo.Update(ctx, doctor); err != nil {
		return errorx.ErrDatabaseOperation
	}

	return nil
}

func (s *DoctorService) Delete(ctx context.Context, id int64) error {
	if err := s.doctorRepo.Delete(ctx, id); err != nil {
		return errorx.ErrDatabaseOperation
	}
	return nil
}
