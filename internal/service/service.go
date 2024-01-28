package service

import (
	"context"
	"github.com/PerfilievAlexandr/auth/internal/api/http/dtoHttpUser"
	"github.com/PerfilievAlexandr/auth/internal/domain"
)

type UserService interface {
	Create(ctx context.Context, req dtoHttpUser.SignUpRequest) (int64, error)
	Get(ctx context.Context, userId int64) (*domain.User, error)
	GetAll(ctx context.Context) ([]*domain.User, error)
	Update(ctx context.Context, userId int64, req dtoHttpUser.UpdateRequest) error
	Delete(ctx context.Context, userId int64) error
}

type PasswordService interface {
	HashAndSaltPassword(ctx context.Context, password string) (string, error)
	CompareWithHashedPassword(ctx context.Context, newPassword string, hashedPassword string) bool
	CompareWithConfirmPassword(_ context.Context, password string, confirmPassword string) bool
}
