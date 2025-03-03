package queries

import (
	"context"
	"shift-scheduling-V2/internal/common/decorator"
	"shift-scheduling-V2/internal/common/utils"
	dto "shift-scheduling-V2/internal/domain/dto/crag"
	"shift-scheduling-V2/internal/domain/entities/crag"
	"shift-scheduling-V2/pkg/logger"
)

// GetAllCragsRequestHandler Contains the dependencies of the Handler
type GetAllCragsRequestHandler decorator.QueryHandler[dto.GetAllCragRequest, []dto.GetAllCragsResult]

type getAllCragsRequestHandler struct {
	repo crag.Repository
}

// NewGetAllCragsRequestHandler Handler constructor
func NewGetAllCragsRequestHandler(repo crag.Repository, logger logger.Logger,
	metricsClient decorator.MetricsClient) GetAllCragsRequestHandler {
	return decorator.ApplyQueryDecorators[dto.GetAllCragRequest, []dto.GetAllCragsResult](
		getAllCragsRequestHandler{repo: repo},
		logger,
		metricsClient,
	)
}

// Handle Handles the query
func (h getAllCragsRequestHandler) Handle(ctx context.Context, _ dto.GetAllCragRequest) ([]dto.GetAllCragsResult, error) {
	res, err := h.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dto.GetAllCragsResult
	for _, modelCrag := range res {
		var cragResult dto.GetAllCragsResult
		err = utils.BindingStruct(modelCrag, &cragResult)
		if err != nil {
			return result, err
		}
		result = append(result, cragResult)
	}
	return result, nil
}
