package config

import (
	dbConfig "auth/internal/config/db"
	grpcConfig "auth/internal/config/grpc"
	configInterface "auth/internal/config/interface"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	GRPCConfig configInterface.GrpcServerConfig
	DbConfig   configInterface.DatabaseConfig
}

func NewConfig() (*Config, error) {
	dbCfg, err := dbConfig.NewDbConfig()
	if err != nil {
		log.Fatalf("failed to config: %s", err.Error())
	}
	grpcCfg, err := grpcConfig.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to config: %s", err.Error())
	}

	return &Config{
		DbConfig:   dbCfg,
		GRPCConfig: grpcCfg,
	}, nil
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
