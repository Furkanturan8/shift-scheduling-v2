package commands

import (
	"context"
	"fmt"
	"shift-scheduling-V2/internal/common/decorator"
	dto "shift-scheduling-V2/internal/domain/dto/crag"
	"shift-scheduling-V2/internal/domain/entities/crag"
	"shift-scheduling-V2/pkg/logger"
)

// UpdateCragRequestHandler Contains the dependencies of the handler
type UpdateCragRequestHandler decorator.CommandHandler[*dto.UpdateCragRequest]

type updateCragRequestHandler struct {
	repo crag.Repository
}

// NewUpdateCragRequestHandler Constructor
func NewUpdateCragRequestHandler(
	repo crag.Repository,
	logger logger.Logger,
	metricsClient decorator.MetricsClient) UpdateCragRequestHandler {

	return decorator.ApplyCommandDecorators[*dto.UpdateCragRequest](
		updateCragRequestHandler{repo: repo},
		logger,
		metricsClient,
	)
}

// Handle Handles the update request
func (h updateCragRequestHandler) Handle(ctx context.Context, command *dto.UpdateCragRequest) error {
	crag, err := h.repo.GetByID(command.ID)
	if crag == nil {
		return fmt.Errorf("the provided crag id does not exist")
	}
	if err != nil {
		return err
	}

	crag.Name = command.Name
	crag.Desc = command.Desc
	crag.Country = command.Country

	return h.repo.Update(*crag)

}
