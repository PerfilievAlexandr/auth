package repository

import (
	"context"
	"github.com/PerfilievAlexandr/auth/internal/api/grpc/user/dto"
	"github.com/PerfilievAlexandr/auth/internal/domain"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserRepository interface {
	Create(ctx context.Context, req *dto.CreateRequest) (int64, error)
	Get(ctx context.Context, id int64) (*domain.User, error)
	Update(ctx context.Context, req *dto.UpdateRequest) (*emptypb.Empty, error)
	Delete(ctx context.Context, id int64) (*emptypb.Empty, error)
}
