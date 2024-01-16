package user

import (
	"auth/internal/api/grpc/user/dto"
	"auth/internal/domain"
	"auth/internal/repository"
	"auth/internal/service"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &userService{
		userRepository: userRepository,
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
