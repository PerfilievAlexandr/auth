package service

import (
	"context"
	"github.com/PerfilievAlexandr/auth/internal/api/grpc/auth/dtoGrpcAuth"
	"github.com/PerfilievAlexandr/auth/internal/api/http/dtoHttpUser"
	"github.com/PerfilievAlexandr/auth/internal/domain"
	"github.com/PerfilievAlexandr/auth/internal/dto"
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
	CompareWithHashedPassword(ctx context.Context, dbPassword string, newPassword string) bool
	CompareWithConfirmPassword(_ context.Context, password string, confirmPassword string) bool
}

type AuthService interface {
	Login(ctx context.Context, req authGrpcDto.LoginRequest) (string, error)
	GetRefreshToken(ctx context.Context, req string) (string, error)
	GetAccessToken(ctx context.Context, req string) (string, error)
}

type JwtService interface {
	GenerateRefreshToken(ctx context.Context, info dto.JwtUserInfo) (string, error)
	GenerateAccessToken(ctx context.Context, info dto.JwtUserInfo) (string, error)
	VerifyRefreshToken(ctx context.Context, token string) (*dto.JwtClaims, error)
	VerifyAccessToken(ctx context.Context, token string) (*dto.JwtClaims, error)
}

type AccessService interface {
	Check(ctx context.Context, endpointAddress string) error
}
