package tests

import (
	"context"
	"fmt"
	"github.com/PerfilievAlexandr/auth/internal/api/grpc/auth"
	authGrpcDto "github.com/PerfilievAlexandr/auth/internal/api/grpc/auth/dtoGrpcAuth"
	"github.com/PerfilievAlexandr/auth/internal/service"
	"github.com/PerfilievAlexandr/auth/internal/service/mocks"
	proto "github.com/PerfilievAlexandr/auth/pkg/auth_v1"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogin(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *proto.LoginRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		name     = "Bob"
		token    = "token123"
		password = "123"

		serviceErr = fmt.Errorf("service error")

		req = &proto.LoginRequest{
			Username: name,
			Password: password,
		}

		mappedReq = authGrpcDto.LoginRequest{
			Username: name,
			Password: password,
		}

		res = &proto.LoginResponse{
			RefreshToken: token,
		}
	)

	t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *proto.LoginResponse
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := mocks.NewAuthServiceMock(mc)
				mock.LoginMock.Expect(ctx, mappedReq).Return(token, nil)
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
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := mocks.NewAuthServiceMock(mc)
				mock.LoginMock.Expect(ctx, mappedReq).Return("", serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			noteServiceMock := tt.authServiceMock(mc)
			api := auth.NewImplementation(noteServiceMock)

			newID, err := api.Login(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
