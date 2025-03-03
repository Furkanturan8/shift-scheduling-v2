package queries

import (
	"context"
	"shift-scheduling-V2/internal/common/decorator"
	"shift-scheduling-V2/internal/common/utils"
	dto "shift-scheduling-V2/internal/domain/dto/user"
	"shift-scheduling-V2/internal/domain/entities/user"
	"shift-scheduling-V2/pkg/logger"
)

type GetUserRequestHandler decorator.QueryHandler[*dto.GetUserRequest, *dto.GetUserResult]

type getUserRequestHandler struct {
	repo user.Repository
}

// NewGetUserRequestHandler Handler Constructor
func NewGetUserRequestHandler(
	repo user.Repository,
	logger logger.Logger,
	metricsClient decorator.MetricsClient) GetUserRequestHandler {
	return decorator.ApplyQueryDecorators[*dto.GetUserRequest, *dto.GetUserResult](
		getUserRequestHandler{repo: repo},
		logger,
		metricsClient,
	)
}

// Handle Handlers the GetUserRequest query
func (h getUserRequestHandler) Handle(ctx context.Context, query *dto.GetUserRequest) (*dto.GetUserResult, error) {
	var result dto.GetUserResult

	userData, err := h.repo.GetByID(query.UserID)
	if err != nil {
		return &result, err
	}
	err = utils.BindingStruct(userData, &result)
	if err != nil {
		return &result, err
	}
	return &result, nil
}
