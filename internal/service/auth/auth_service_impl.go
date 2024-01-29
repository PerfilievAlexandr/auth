package authService

import (
	"context"
	"errors"
	"github.com/PerfilievAlexandr/auth/internal/api/grpc/auth/dtoGrpcAuth"
	"github.com/PerfilievAlexandr/auth/internal/dto"
	"github.com/PerfilievAlexandr/auth/internal/repository"
	"github.com/PerfilievAlexandr/auth/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authService struct {
	userRepository  repository.UserRepository
	passwordService service.PasswordService
	jwtService      service.JwtService
}

func NewAuthService(
	userRepository repository.UserRepository,
	passwordService service.PasswordService,
	jwtService service.JwtService,
) service.AuthService {
	return &authService{
		userRepository:  userRepository,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

func (a *authService) Login(ctx context.Context, req authGrpcDto.LoginRequest) (string, error) {
	user, err := a.userRepository.GetByUsername(ctx, req.Username)
	if err != nil {
		return "", errors.New("user not found")
	}

	isPasswordsEquals := a.passwordService.CompareWithHashedPassword(ctx, user.Password, req.Password)
	if !isPasswordsEquals {
		return "", errors.New("wrong password")
	}

	token, err := a.jwtService.GenerateRefreshToken(
		ctx,
		dto.JwtUserInfo{
			Username: user.Name,
			Role:     user.Role,
		},
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (a *authService) GetRefreshToken(ctx context.Context, oldToken string) (string, error) {
	claims, err := a.jwtService.VerifyRefreshToken(ctx, oldToken)
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	token, err := a.jwtService.GenerateRefreshToken(
		ctx,
		dto.JwtUserInfo{
			Username: claims.Username,
			Role:     claims.Role,
		},
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (a *authService) GetAccessToken(ctx context.Context, oldToken string) (string, error) {
	claims, err := a.jwtService.VerifyRefreshToken(ctx, oldToken)
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	token, err := a.jwtService.GenerateAccessToken(
		ctx,
		dto.JwtUserInfo{
			Username: claims.Username,
			Role:     claims.Role,
		},
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
