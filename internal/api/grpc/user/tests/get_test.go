package tests

import (
	"auth/internal/api/grpc/user"
	"auth/internal/domain"
	"auth/internal/service"
	serviceMocks "auth/internal/service/mocks"
	proto "auth/pkg/user_v1"
	"context"
	"database/sql"
	"fmt"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *proto.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = int64(7)
		name      = "Bob"
		email     = "bob@gmail.ru"
		createdAt = time.Now()
		role      = "user"

		serviceErr = fmt.Errorf("service error")

		req = &proto.GetRequest{
			Id: id,
		}

		nullTime = sql.NullTime{
			Valid: false,
		}

		res = &proto.GetResponse{
			Id:        id,
			Name:      name,
			Email:     email,
			Role:      role,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: nil,
		}

		userDomain = &domain.User{
			Id:        id,
			Name:      name,
			Email:     email,
			Role:      role,
			CreatedAt: createdAt,
			UpdatedAt: nullTime,
		}
	)

	t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *proto.GetResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(userDomain, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			noteServiceMock := tt.userServiceMock(mc)
			api := user.NewImplementation(noteServiceMock)

			newID, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
