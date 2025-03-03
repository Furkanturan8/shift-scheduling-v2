package commands

import (
	"context"
	"shift-scheduling-V2/internal/common/decorator"
	dto "shift-scheduling-V2/internal/domain/dto/user"
	"shift-scheduling-V2/internal/domain/entities/notification"
	user "shift-scheduling-V2/internal/domain/entities/user"
	"shift-scheduling-V2/pkg/logger"
	timePkg "shift-scheduling-V2/pkg/time"
	uuidPkg "shift-scheduling-V2/pkg/uuid"
)

type AddUserRequestHandler decorator.CommandHandler[*dto.AddUserRequest]

type addUserRequestHandler struct {
	uuidProvider        uuidPkg.Provider
	timeProvider        timePkg.Provider
	repo                user.Repository
	notificationService notification.Service
}

// NewAddUserRequestHandler Initializes an AddCommandHandler
func NewAddUserRequestHandler(
	uuidProvider uuidPkg.Provider,
	timeProvider timePkg.Provider,
	repo user.Repository,
	notificationService notification.Service,
	logger logger.Logger,
	metricsClient decorator.MetricsClient) AddUserRequestHandler {
	return decorator.ApplyCommandDecorators[*dto.AddUserRequest](
		addUserRequestHandler{uuidProvider: uuidProvider, timeProvider: timeProvider, repo: repo, notificationService: notificationService},
		logger,
		metricsClient,
	)
}

// Handle Handles the AddUserRequest
func (h addUserRequestHandler) Handle(ctx context.Context, req *dto.AddUserRequest) error {
	c := user.User{
		ID:        h.uuidProvider.NewUUID(),
		Name:      req.Name,
		Desc:      req.Desc,
		Country:   req.Country,
		CreatedAt: h.timeProvider.Now(),
	}
	err := h.repo.Add(c)
	if err != nil {
		return err
	}
	n := notification.Notification{
		Subject: "New User added",
		Message: "A new User with name '" + c.Name + "' was added in the repository",
	}
	return h.notificationService.Notify(n)
}
