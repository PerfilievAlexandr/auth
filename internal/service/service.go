package service

import (
	"auth/internal/api/grpc/user/dto"
	"auth/internal/domain"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService interface {
	Create(ctx context.Context, req *dto.CreateRequest) (int64, error)
	Get(ctx context.Context, id int64) (*domain.User, error)
	Update(ctx context.Context, req *dto.UpdateRequest) (*emptypb.Empty, error)
	Delete(ctx context.Context, id int64) (*emptypb.Empty, error)
}
