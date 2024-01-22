package config

import (
	"context"
	dbConfig "github.com/PerfilievAlexandr/auth/internal/config/db"
	grpcConfig "github.com/PerfilievAlexandr/auth/internal/config/grpc"
	httpConfig "github.com/PerfilievAlexandr/auth/internal/config/http"
	configInterface "github.com/PerfilievAlexandr/auth/internal/config/interface"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	HttpConfig configInterface.HttpServerConfig
	GRPCConfig configInterface.GrpcServerConfig
	DbConfig   configInterface.DatabaseConfig
}

func NewConfig(_ context.Context) (*Config, error) {
	dbCfg, err := dbConfig.NewDbConfig()
	if err != nil {
		log.Fatalf("failed to config: %s", err.Error())
	}
	grpcCfg, err := grpcConfig.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to config: %s", err.Error())
	}
	httpCfg, err := httpConfig.NewHttpConfig()
	if err != nil {
		log.Fatalf("failed to config: %s", err.Error())
	}

	return &Config{
		DbConfig:   dbCfg,
		GRPCConfig: grpcCfg,
		HttpConfig: httpCfg,
	}, nil
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
