package config

import (
	configInterface "github.com/PerfilievAlexandr/auth/internal/config/interface"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

var _ configInterface.JwtConfig = (*jwtConfig)(nil)

const (
	refreshTokenSecret     = "REFRESH_TOKEN_SECRET"
	accessTokenSecret      = "ACCESS_TOKEN_SECRET"
	refreshTokenExpiration = "REFRESH_TOKEN_EXPIRATION"
	accessTokenExpiration  = "ACCESS_TOKEN_EXPIRATION"
)

type jwtConfig struct {
	refreshTokenSecret     string
	accessTokenSecret      string
	refreshTokenExpiration time.Duration
	accessTokenExpiration  time.Duration
}

func NewJwtConfig() (configInterface.JwtConfig, error) {
	refreshSecret := os.Getenv(refreshTokenSecret)
	if len(refreshSecret) == 0 {
		return nil, errors.New("refresh secret not found")
	}

	accessSecret := os.Getenv(accessTokenSecret)
	if len(accessSecret) == 0 {
		return nil, errors.New("access secret not found")
	}

	refreshExpiration := os.Getenv(refreshTokenExpiration)
	if len(refreshExpiration) == 0 {
		return nil, errors.New("refresh expiration not found")
	}
	refreshExpirationTime, err := strconv.Atoi(refreshExpiration)
	if err != nil {
		return nil, errors.New("refreshTokenExpiration is not a number")
	}

	accessExpiration := os.Getenv(accessTokenExpiration)
	if len(accessExpiration) == 0 {
		return nil, errors.New("access expiration not found")
	}
	accessExpirationTime, err := strconv.Atoi(accessExpiration)
	if err != nil {
		return nil, errors.New("accessTokenExpiration is not a number")
	}

	return &jwtConfig{
		refreshTokenSecret:     refreshSecret,
		accessTokenSecret:      accessSecret,
		refreshTokenExpiration: time.Duration(refreshExpirationTime),
		accessTokenExpiration:  time.Duration(accessExpirationTime),
	}, nil
}

func (j jwtConfig) RefreshTokenSecret() string {
	return j.refreshTokenSecret
}

func (j jwtConfig) AccessTokenSecret() string {
	return j.accessTokenSecret
}

func (j jwtConfig) RefreshTokenExpirationMinutes() time.Duration {
	return j.refreshTokenExpiration * time.Minute
}

func (j jwtConfig) AccessTokenExpirationMinutes() time.Duration {
	return j.accessTokenExpiration * time.Minute
}
