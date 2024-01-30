package config

import (
	"context"
	dbConfig "github.com/PerfilievAlexandr/auth/internal/config/db"
	grpcConfig "github.com/PerfilievAlexandr/auth/internal/config/grpc"
	httpConfig "github.com/PerfilievAlexandr/auth/internal/config/http"
	configInterface "github.com/PerfilievAlexandr/auth/internal/config/interface"
	jwtConfig "github.com/PerfilievAlexandr/auth/internal/config/jwt"
	"github.com/PerfilievAlexandr/auth/internal/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	HttpConfig configInterface.HttpServerConfig
	GRPCConfig configInterface.GrpcServerConfig
	DbConfig   configInterface.DatabaseConfig
	JwtConfig  configInterface.JwtConfig
}

func NewConfig(_ context.Context) (*Config, error) {
	dbCfg, err := dbConfig.NewDbConfig()
	if err != nil {
		logger.Fatal("failed to config", zap.Any("err", err))
	}
	grpcCfg, err := grpcConfig.NewGRPCConfig()
	if err != nil {
		logger.Fatal("failed to config", zap.Any("err", err))
	}
	httpCfg, err := httpConfig.NewHttpConfig()
	if err != nil {
		logger.Fatal("failed to config", zap.Any("err", err))
	}
	jwtCfg, err := jwtConfig.NewJwtConfig()
	if err != nil {
		logger.Fatal("failed to config", zap.Any("err", err))
	}

	return &Config{
		DbConfig:   dbCfg,
		GRPCConfig: grpcCfg,
		HttpConfig: httpCfg,
		JwtConfig:  jwtCfg,
	}, nil
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
