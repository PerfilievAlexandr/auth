package jwtService

import (
	"context"
	"github.com/PerfilievAlexandr/auth/internal/config"
	"github.com/PerfilievAlexandr/auth/internal/dto"
	"github.com/PerfilievAlexandr/auth/internal/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

type jwtService struct {
	config *config.Config
}

func NewJwtService(
	config *config.Config,
) service.JwtService {
	return &jwtService{
		config: config,
	}
}

func (j *jwtService) GenerateRefreshToken(ctx context.Context, info dto.JwtUserInfo) (string, error) {
	return j.generateToken(ctx, info, j.config.JwtConfig.RefreshTokenExpirationMinutes(), j.config.JwtConfig.RefreshTokenSecret())
}

func (j *jwtService) GenerateAccessToken(ctx context.Context, info dto.JwtUserInfo) (string, error) {
	return j.generateToken(ctx, info, j.config.JwtConfig.AccessTokenExpirationMinutes(), j.config.JwtConfig.AccessTokenSecret())
}

func (j *jwtService) generateToken(_ context.Context, info dto.JwtUserInfo, expireTime time.Duration, secret string) (string, error) {
	claims := dto.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireTime).Unix(),
		},
		Username: info.Username,
		Role:     info.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (j *jwtService) VerifyRefreshToken(ctx context.Context, token string) (*dto.JwtClaims, error) {
	return j.verifyToken(ctx, token, j.config.JwtConfig.RefreshTokenSecret())
}

func (j *jwtService) VerifyAccessToken(ctx context.Context, token string) (*dto.JwtClaims, error) {
	return j.verifyToken(ctx, token, j.config.JwtConfig.AccessTokenSecret())
}

func (j *jwtService) verifyToken(_ context.Context, tokenStr string, secret string) (*dto.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&dto.JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.Errorf("unexpected token signing method")
			}

			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, errors.Errorf("invalid token: %s", err.Error())
	}

	claims, ok := token.Claims.(*dto.JwtClaims)
	if !ok {
		return nil, errors.Errorf("invalid token claims")
	}

	return claims, nil
}
