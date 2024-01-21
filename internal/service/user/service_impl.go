package user

import (
	"context"
	"github.com/PerfilievAlexandr/auth/internal/api/grpc/user/dto"
	"github.com/PerfilievAlexandr/auth/internal/domain"
	"github.com/PerfilievAlexandr/auth/internal/repository"
	"github.com/PerfilievAlexandr/auth/internal/service"
	"github.com/PerfilievAlexandr/platform_common/pkg/db"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userService struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewUserService(userRepository repository.UserRepository, txManager db.TxManager) service.UserService {
	return &userService{
		userRepository: userRepository,
		txManager:      txManager,
	}
}

func (s *userService) Create(ctx context.Context, req *dto.CreateRequest) (int64, error) {
	return s.userRepository.Create(ctx, req)
}

func (s *userService) Get(ctx context.Context, userId int64) (*domain.User, error) {
	return s.userRepository.Get(ctx, userId)
}

func (s *userService) Update(ctx context.Context, req *dto.UpdateRequest) (*emptypb.Empty, error) {
	return s.userRepository.Update(ctx, req)
}

func (s *userService) Delete(ctx context.Context, userId int64) (*emptypb.Empty, error) {
	return s.userRepository.Delete(ctx, userId)
}
