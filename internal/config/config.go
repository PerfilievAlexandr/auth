package config

import (
	db1Config "auth/internal/config/db"
	grpc1Config "auth/internal/config/grpc"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	GRPCConfig grpcConfig
	DbConfig   dbConfig
}

func NewConfig() (*Config, error) {
	dbCfg, err := db1Config.NewDbConfig()
	if err != nil {
		log.Fatalf("failed to config: %s", err.Error())
	}
	grpcCfg, err := grpc1Config.NewGRPCConfig()
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

type grpcConfig interface {
	Address() string
}

type dbConfig interface {
	ConnectString() string
}
