package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"shift-scheduling-V2/internal/app"
	"shift-scheduling-V2/internal/common/errors"
	"shift-scheduling-V2/internal/common/responses"
	"shift-scheduling-V2/internal/common/validator"
	dto "shift-scheduling-V2/internal/domain/dto/user"
)

type UserHttpApi interface {
	AddUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
	GetUsers(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
}

type userHttpApi struct {
	userApp app.Application
}

// NewHandler Constructor
func NewUserHttpApi(userApp app.Application) UserHttpApi {
	return &userHttpApi{userApp: userApp}
}

// GetUser GetById swagger documentation
// @Summary Get a User by ID
// @Description Get a User by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Failure 400 {object} responses.General
// @Router /User/{id} [get]
func (cr *userHttpApi) GetUser(ctx *fiber.Ctx) error {
	context := ctx.Context()

	_UserId := ctx.Params("id", "")
	if _UserId == "" {
		return errors.ErrBadRequest
	}

	UserId, err := uuid.Parse(_UserId)
	if err != nil {
		return errors.ErrBadRequest
	}

	req := &dto.GetUserRequest{UserID: UserId}

	User, err := cr.userApp.Queries.GetUserHandler.Handle(context, req)
	if err != nil {
		return err
	}

	resp := responses.DefaultSuccessResponse
	resp.Data = User
	return resp.JSON(ctx)
}

// GetUsers GetAll Returns all available Users
// @Summary Get all Users
// @Tags User
// @Produce json
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Failure 400 {object} responses.General
// @Router /User [get]
func (cr *userHttpApi) GetUsers(ctx *fiber.Ctx) error {
	Users, err := cr.userApp.Queries.GetAllUsersHandler.Handle(ctx.Context(), dto.GetAllUserRequest{})
	if err != nil {
		return err
	}
	resp := responses.DefaultSuccessResponse
	resp.Data = Users
	return resp.JSON(ctx)
}

// AddUser Add a new User
// @Summary Add a new User
// @Tags User
// @Accept json
// @Produce json
// @Param User body dto.AddUserRequest true "The User data"
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Failure 400 {object} responses.General
// @Router /User [post]
func (cr *userHttpApi) AddUser(ctx *fiber.Ctx) error {
	context := ctx.Context()

	req := new(dto.AddUserRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return errors.ErrBadRequest
	}

	if err := validator.GetValidator().Validate(req); err != nil {
		return errors.ErrBadRequest
	}

	if err := cr.userApp.Commands.AddUserHandler.Handle(context, req); err != nil {
		return err
	}
	return responses.DefaultSuccessResponse.JSON(ctx)
}

// UpdateUser swagger documentation
// @Summary Update a User
// @Description Update a User by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body dto.UpdateUserRequest true "UpdateUserRequest object"
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Failure 400 {object} responses.General
// @Router /User/{id} [put]
func (cr *userHttpApi) UpdateUser(ctx *fiber.Ctx) error {
	context := ctx.Context()

	_UserId := ctx.Params("id", "")
	if _UserId == "" {
		return errors.ErrBadRequest
	}

	UserId, err := uuid.Parse(_UserId)
	if err != nil {
		return errors.ErrBadRequest
	}

	req := new(dto.UpdateUserRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return errors.ErrBadRequest
	}

	if err := validator.GetValidator().Validate(req); err != nil {
		return errors.ErrBadRequest
	}
	req.ID = UserId

	if err := cr.userApp.Commands.UpdateUserHandler.Handle(context, req); err != nil {
		return err
	}
	return responses.DefaultSuccessResponse.JSON(ctx)
}

// DeleteUser swagger documentation
// @Summary Delete a User
// @Description Delete a User by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Failure 400 {object} responses.General
// @Router /User/{id} [delete]
func (cr *userHttpApi) DeleteUser(ctx *fiber.Ctx) error {
	context := ctx.Context()

	_UserId := ctx.Params("id", "")
	if _UserId == "" {
		return errors.ErrBadRequest
	}

	UserId, err := uuid.Parse(_UserId)
	if err != nil {
		return errors.ErrBadRequest
	}

	req := &dto.DeleteUserRequest{
		UserID: UserId,
	}

	if err := cr.userApp.Commands.DeleteUserHandler.Handle(context, req); err != nil {
		return err
	}
	return responses.DefaultSuccessResponse.JSON(ctx)
}
