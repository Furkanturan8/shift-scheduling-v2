package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	dto "shift-scheduling-V2/internal/domain/dto/user"

	"shift-scheduling-V2/internal/domain/entities/user"
	"testing"
	"time"
)

func TestUpdateUserCommandHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		repo user.Repository
	}
	type args struct {
		command *dto.UpdateUserRequest
		ctx     context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		err    error
	}{
		{
			name: "happy path - no errors - should return nil",
			fields: fields{
				repo: func() user.MockRepository {
					mp := user.MockRepository{}
					returnedUser := user.User{
						ID:        mockUUID,
						Name:      "initial",
						Surname:   "initial",
						UserName:  "initial",
						Email:     "initial",
						Password:  "initial",
						CreatedAt: time.Time{},
					}
					updatedUser := user.User{
						ID:        mockUUID,
						Name:      "updated",
						Surname:   "updated",
						UserName:  "updated",
						Email:     "updated",
						Password:  "updated",
						CreatedAt: time.Time{},
					}
					mp.On("GetByID", mockUUID).Return(&returnedUser, nil)
					mp.On("Update", updatedUser).Return(nil)

					return mp
				}(),
			},
			args: args{
				command: &dto.UpdateUserRequest{
					ID:       mockUUID,
					Name:     "updated",
					Surname:  "updated",
					UserName: "updated",
					Email:    "updated",
					Password: "updated",
				},
				ctx: context.Background(),
			},
			err: nil,
		},
		{
			name: "get error should return error",
			fields: fields{
				repo: func() user.MockRepository {
					mp := user.MockRepository{}
					mp.On("GetByID", mockUUID).Return(&user.User{ID: mockUUID}, errors.New("get error"))

					return mp
				}(),
			},
			args: args{
				command: &dto.UpdateUserRequest{
					ID:   mockUUID,
					Name: "updated",
				},
				ctx: context.Background(),
			},
			err: errors.New("get error"),
		},
		{
			name: "get returns nil, should return error",
			fields: fields{
				repo: func() user.MockRepository {
					mp := user.MockRepository{}
					mp.On("GetByID", mockUUID).Return((*user.User)(nil), nil)
					return mp
				}(),
			},
			args: args{
				command: &dto.UpdateUserRequest{
					ID:   mockUUID,
					Name: "updated",
				},
				ctx: context.Background(),
			},
			err: fmt.Errorf("the provided crag id does not exist"),
		},
		{
			name: "update error - should return error",
			fields: fields{
				repo: func() user.MockRepository {
					mp := user.MockRepository{}
					returnedUser := user.User{
						ID:        mockUUID,
						Name:      "initial",
						CreatedAt: time.Time{},
					}
					updatedUser := user.User{
						ID:        mockUUID,
						Name:      "updated",
						CreatedAt: time.Time{},
					}
					mp.On("GetByID", mockUUID).Return(&returnedUser, nil)
					mp.On("Update", updatedUser).Return(errors.New("update error"))

					return mp
				}(),
			},
			args: args{
				command: &dto.UpdateUserRequest{
					ID:   mockUUID,
					Name: "updated",
				},
				ctx: context.Background(),
			},
			err: errors.New("update error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := updateUserRequestHandler{
				repo: tt.fields.repo,
			}
			err := h.Handle(tt.args.ctx, tt.args.command)
			assert.Equal(t, tt.err, err)
		})
	}
}
