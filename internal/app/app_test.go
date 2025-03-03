package app

import (
	"shift-scheduling-V2/internal/app/crag/commands"
	"shift-scheduling-V2/internal/app/crag/queries"
	userCommands "shift-scheduling-V2/internal/app/user/commands"
	userQuery "shift-scheduling-V2/internal/app/user/queries"
	"shift-scheduling-V2/internal/common/metrics"
	"shift-scheduling-V2/internal/domain/entities/crag"
	"shift-scheduling-V2/internal/domain/entities/notification"
	user "shift-scheduling-V2/internal/domain/entities/user"
	logger2 "shift-scheduling-V2/pkg/logger"
	"shift-scheduling-V2/pkg/time"
	"shift-scheduling-V2/pkg/uuid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	cragMockRepo := crag.MockRepository{}
	userMockRepo := user.MockRepository{}
	notificationService := notification.MockNotificationService{}
	// init base
	logger := logger2.NewApiLogger()
	metricsClient := metrics.NoOp{}
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()
	type args struct {
		cragRepo            crag.Repository
		userRepo            user.Repository
		notificationService notification.Service
	}
	tests := []struct {
		name string
		args args
		want Application
	}{
		{
			name: "should initialize application layer",
			args: args{
				cragRepo:            cragMockRepo,
				userRepo:            userMockRepo,
				notificationService: notificationService,
			},
			want: Application{
				Queries: Queries{
					GetAllCragsHandler: queries.NewGetAllCragsRequestHandler(cragMockRepo, logger, metricsClient),
					GetCragHandler:     queries.NewGetCragRequestHandler(cragMockRepo, logger, metricsClient),
					GetUserHandler:     userQuery.NewGetUserRequestHandler(userMockRepo, logger, metricsClient),
					GetAllUsersHandler: userQuery.NewGetAllUsersRequestHandler(userMockRepo, logger, metricsClient),
				},
				Commands: Commands{
					AddCragHandler:    commands.NewAddCragRequestHandler(up, tp, cragMockRepo, notificationService, logger, metricsClient),
					UpdateCragHandler: commands.NewUpdateCragRequestHandler(cragMockRepo, logger, metricsClient),
					DeleteCragHandler: commands.NewDeleteCragRequestHandler(cragMockRepo, logger, metricsClient),
					AddUserHandler:    userCommands.NewAddUserRequestHandler(up, tp, userMockRepo, notificationService, logger, metricsClient),
					UpdateUserHandler: userCommands.NewUpdateUserRequestHandler(userMockRepo, logger, metricsClient),
					DeleteUserHandler: userCommands.NewDeleteUserRequestHandler(userMockRepo, logger, metricsClient),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewApplication(tt.args.cragRepo, tt.args.userRepo, tt.args.notificationService, logger)
			assert.Equal(t, tt.want, got)
		})
	}
}
