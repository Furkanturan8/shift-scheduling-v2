package app

import (
	cragCommands "shift-scheduling-V2/internal/app/crag/commands"
	cragQuery "shift-scheduling-V2/internal/app/crag/queries"
	userCommands "shift-scheduling-V2/internal/app/user/commands"
	userQuery "shift-scheduling-V2/internal/app/user/queries"
	"shift-scheduling-V2/internal/common/metrics"
	"shift-scheduling-V2/internal/domain/entities/crag"
	"shift-scheduling-V2/internal/domain/entities/notification"
	"shift-scheduling-V2/internal/domain/entities/user"
	"shift-scheduling-V2/pkg/logger"
	"shift-scheduling-V2/pkg/time"
	"shift-scheduling-V2/pkg/uuid"
)

// Queries Contains all available query handlers of this app
type Queries struct {
	GetAllCragsHandler cragQuery.GetAllCragsRequestHandler
	GetCragHandler     cragQuery.GetCragRequestHandler
	GetUserHandler     userQuery.GetUserRequestHandler
	GetAllUsersHandler userQuery.GetAllUsersRequestHandler
}

// Commands Contains all available command handlers of this app
type Commands struct {
	AddCragHandler    cragCommands.AddCragRequestHandler
	UpdateCragHandler cragCommands.UpdateCragRequestHandler
	DeleteCragHandler cragCommands.DeleteCragRequestHandler
	AddUserHandler    userCommands.AddUserRequestHandler
	UpdateUserHandler userCommands.UpdateUserRequestHandler
	DeleteUserHandler userCommands.DeleteUserRequestHandler
}

type Application struct {
	Queries  Queries
	Commands Commands
}

func NewApplication(cragRepo crag.Repository, userRepo user.Repository, ns notification.Service, logger logger.Logger) Application {
	// init base
	metricsClient := metrics.NoOp{}
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()
	return Application{
		Queries: Queries{
			GetAllCragsHandler: cragQuery.NewGetAllCragsRequestHandler(cragRepo, logger, metricsClient),
			GetCragHandler:     cragQuery.NewGetCragRequestHandler(cragRepo, logger, metricsClient),
			GetUserHandler:     userQuery.NewGetUserRequestHandler(userRepo, logger, metricsClient),
			GetAllUsersHandler: userQuery.NewGetAllUsersRequestHandler(userRepo, logger, metricsClient),
		},
		Commands: Commands{
			AddCragHandler:    cragCommands.NewAddCragRequestHandler(up, tp, cragRepo, ns, logger, metricsClient),
			UpdateCragHandler: cragCommands.NewUpdateCragRequestHandler(cragRepo, logger, metricsClient),
			DeleteCragHandler: cragCommands.NewDeleteCragRequestHandler(cragRepo, logger, metricsClient),
		},
	}
}
