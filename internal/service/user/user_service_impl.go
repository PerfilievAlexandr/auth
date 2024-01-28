package userService

import (
	"context"
	"errors"
	"github.com/PerfilievAlexandr/auth/internal/api/http/dtoHttpUser"
	"github.com/PerfilievAlexandr/auth/internal/domain"
	"github.com/PerfilievAlexandr/auth/internal/dto"
	"github.com/PerfilievAlexandr/auth/internal/repository"
	"github.com/PerfilievAlexandr/auth/internal/service"
	"github.com/PerfilievAlexandr/platform_common/pkg/db"
)

type userService struct {
	userRepository  repository.UserRepository
	passwordService service.PasswordService
	txManager       db.TxManager
}

func NewUserService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
	passwordService service.PasswordService,
) service.UserService {
	return &userService{
		userRepository:  userRepository,
		txManager:       txManager,
		passwordService: passwordService,
	}
}

func (s *userService) Create(ctx context.Context, req dtoHttpUser.SignUpRequest) (int64, error) {
	isPasswordEquals := s.passwordService.CompareWithConfirmPassword(ctx, req.Password, req.PasswordConfirm)
	if !isPasswordEquals {
		return 0, errors.New("passwords aren't equals")
	}

	password, err := s.passwordService.HashAndSaltPassword(ctx, req.Password)
	if err != nil {
		return 0, errors.New("hashing error")
	}
	createUser := dto.CreateUser{
		Password: password,
		Name:     req.Name,
		Role:     req.Role,
		Email:    req.Email,
	}

	return s.userRepository.Create(ctx, createUser)
}

func (s *userService) Get(ctx context.Context, userId int64) (*domain.User, error) {
	return s.userRepository.Get(ctx, userId)
}

func (s *userService) GetAll(ctx context.Context) ([]*domain.User, error) {
	return s.userRepository.GetAll(ctx)
}

func (s *userService) Update(ctx context.Context, userId int64, req dtoHttpUser.UpdateRequest) error {
	return s.userRepository.Update(ctx, userId, req)
}

func (s *userService) Delete(ctx context.Context, userId int64) error {
	return s.userRepository.Delete(ctx, userId)
}
