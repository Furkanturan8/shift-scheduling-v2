package service

import (
	"context"
	"shift-scheduling-v2/internal/dto"
	"shift-scheduling-v2/internal/repository"
	"shift-scheduling-v2/pkg/errorx"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) List(ctx context.Context) ([]dto.UserResponseDTO, error) {
	users, err := s.userRepo.List(ctx)
	if err != nil {
		return nil, errorx.ErrDatabaseOperation
	}

	var userList []dto.UserResponseDTO
	for _, user := range users {
		uDto := dto.UserResponseDTO{}.ToResponseModel(user)
		userList = append(userList, uDto)
	}

	return userList, nil
}

func (s *UserService) GetByID(ctx context.Context, id int64) (dto.UserResponseDTO, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return dto.UserResponseDTO{}, errorx.ErrNotFound
	}

	return dto.UserResponseDTO{}.ToResponseModel(*user), nil
}

func (s *UserService) Update(ctx context.Context, id int64, req *dto.UserCreateDTO) error {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return errorx.ErrNotFound
	}

	if user.Email != req.Email {
		exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
		if err != nil {
			return errorx.ErrDatabaseOperation
		}
		if exists {
			return errorx.ErrDuplicate
		}
	}

	req.ToDBModel(*user)

	if err = s.userRepo.Update(ctx, user); err != nil {
		return errorx.ErrDatabaseOperation
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	if err := s.userRepo.Delete(ctx, id); err != nil {
		return errorx.ErrDatabaseOperation
	}
	return nil
}
