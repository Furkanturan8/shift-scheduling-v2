package commands

import (
	"context"
	"fmt"
	"shift-scheduling-V2/internal/common/decorator"
	dto "shift-scheduling-V2/internal/domain/dto/user"
	"shift-scheduling-V2/internal/domain/entities/user"
	"shift-scheduling-V2/pkg/logger"
)

// UpdateUserRequestHandler Contains the dependencies of the handler
type UpdateUserRequestHandler decorator.CommandHandler[*dto.UpdateUserRequest]

type updateUserRequestHandler struct {
	repo user.Repository
}

// NewUpdateUserRequestHandler Constructor
func NewUpdateUserRequestHandler(
	repo user.Repository,
	logger logger.Logger,
	metricsClient decorator.MetricsClient) UpdateUserRequestHandler {

	return decorator.ApplyCommandDecorators[*dto.UpdateUserRequest](
		updateUserRequestHandler{repo: repo},
		logger,
		metricsClient,
	)
}

// Handle Handles the update request
func (h updateUserRequestHandler) Handle(ctx context.Context, command *dto.UpdateUserRequest) error {
	user, err := h.repo.GetByID(command.ID)
	if user == nil {
		return fmt.Errorf("the provided user id does not exist")
	}
	if err != nil {
		return err
	}

	user.Name = command.Name
	user.Desc = command.Desc
	user.Country = command.Country

	return h.repo.Update(*user)

}
