package queries

import (
	"context"
	"shift-scheduling-V2/internal/common/decorator"
	"shift-scheduling-V2/internal/common/utils"
	dto "shift-scheduling-V2/internal/domain/dto/user"
	"shift-scheduling-V2/internal/domain/entities/user"
	"shift-scheduling-V2/pkg/logger"
)

// GetAllUsersRequestHandler Contains the dependencies of the Handler
type GetAllUsersRequestHandler decorator.QueryHandler[dto.GetAllUserRequest, []dto.GetAllUsersResult]

type getAllUsersRequestHandler struct {
	repo user.Repository
}

// NewGetAllUsersRequestHandler Handler constructor
func NewGetAllUsersRequestHandler(repo user.Repository, logger logger.Logger,
	metricsClient decorator.MetricsClient) GetAllUsersRequestHandler {
	return decorator.ApplyQueryDecorators[dto.GetAllUserRequest, []dto.GetAllUsersResult](
		getAllUsersRequestHandler{repo: repo},
		logger,
		metricsClient,
	)
}

// Handle Handles the query
func (h getAllUsersRequestHandler) Handle(ctx context.Context, _ dto.GetAllUserRequest) ([]dto.GetAllUsersResult, error) {
	res, err := h.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dto.GetAllUsersResult
	for _, modelUser := range res {
		var userResult dto.GetAllUsersResult
		err = utils.BindingStruct(modelUser, &userResult)
		if err != nil {
			return result, err
		}
		result = append(result, userResult)
	}
	return result, nil
}
