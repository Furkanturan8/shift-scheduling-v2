package commands

import (
	"context"
	"shift-scheduling-V2/internal/common/decorator"
	dto "shift-scheduling-V2/internal/domain/dto/crag"
	"shift-scheduling-V2/internal/domain/entities/crag"
	"shift-scheduling-V2/internal/domain/entities/notification"
	"shift-scheduling-V2/pkg/logger"

	timePkg "shift-scheduling-V2/pkg/time"
	uuidPkg "shift-scheduling-V2/pkg/uuid"
)

type AddCragRequestHandler decorator.CommandHandler[*dto.AddCragRequest]

type addCragRequestHandler struct {
	uuidProvider        uuidPkg.Provider
	timeProvider        timePkg.Provider
	repo                crag.Repository
	notificationService notification.Service
}

// NewAddCragRequestHandler Initializes an AddCommandHandler
func NewAddCragRequestHandler(
	uuidProvider uuidPkg.Provider,
	timeProvider timePkg.Provider,
	repo crag.Repository,
	notificationService notification.Service,
	logger logger.Logger,
	metricsClient decorator.MetricsClient) AddCragRequestHandler {
	return decorator.ApplyCommandDecorators[*dto.AddCragRequest](
		addCragRequestHandler{uuidProvider: uuidProvider, timeProvider: timeProvider, repo: repo, notificationService: notificationService},
		logger,
		metricsClient,
	)
}

// Handle Handles the AddCragRequest
func (h addCragRequestHandler) Handle(ctx context.Context, req *dto.AddCragRequest) error {
	c := crag.Crag{
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
		Subject: "New crag added",
		Message: "A new crag with name '" + c.Name + "' was added in the repository",
	}
	return h.notificationService.Notify(n)
}
