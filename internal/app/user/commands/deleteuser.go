package commands

import (
	"context"
	"fmt"
	"shift-scheduling-V2/internal/common/decorator"
	dto "shift-scheduling-V2/internal/domain/dto/user"
	"shift-scheduling-V2/internal/domain/entities/user"
	"shift-scheduling-V2/pkg/logger"
)

// DeleteUserRequestHandler Handler Struct with Dependencies
type DeleteUserRequestHandler decorator.CommandHandler[*dto.DeleteUserRequest]

type deleteUserRequestHandler struct {
	repo user.Repository
}

// NewDeleteUserRequestHandler Handler constructor
func NewDeleteUserRequestHandler(
	repo user.Repository,
	logger logger.Logger,
	metricsClient decorator.MetricsClient) DeleteUserRequestHandler {
	return decorator.ApplyCommandDecorators[*dto.DeleteUserRequest](
		deleteUserRequestHandler{repo: repo},
		logger,
		metricsClient)
}

// Handle Handlers the DeleteUserRequest request
func (h deleteUserRequestHandler) Handle(ctx context.Context, command *dto.DeleteUserRequest) error {
	user, err := h.repo.GetByID(command.ID)
	if user == nil {
		return fmt.Errorf("the provided user id does not exist")
	}
	if err != nil {
		return err
	}
	return h.repo.Delete(command.ID)
}
