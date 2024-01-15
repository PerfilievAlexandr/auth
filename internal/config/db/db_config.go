package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

//var _ config.DbConfig = (*DbConfig)(nil)

const (
	port         = "DB_PORT"
	host         = "DB_HOST"
	driver       = "DB_DRIVER"
	user         = "DB_USER"
	password     = "DB_PASSWORD"
	sslMode      = "DB_SSL_MODE"
	databaseName = "DB_NAME"
)

type DbConfig struct {
	Port         string
	Host         string
	Driver       string
	User         string
	Password     string
	SslMode      string
	DatabaseName string
}

func NewDbConfig() (*DbConfig, error) {
	port := os.Getenv(port)
	if len(port) == 0 {
		return nil, errors.New("db Port not found")
	}

	host := os.Getenv(host)
	if len(host) == 0 {
		return nil, errors.New("db Host not found")
	}

	driver := os.Getenv(driver)
	if len(driver) == 0 {
		return nil, errors.New("db Driver not found")
	}

	user := os.Getenv(user)
	if len(user) == 0 {
		return nil, errors.New("db User not found")
	}

	password := os.Getenv(password)
	if len(password) == 0 {
		return nil, errors.New("db Password not found")
	}

	sslMode := os.Getenv(sslMode)
	if len(sslMode) == 0 {
		return nil, errors.New("db SslMode not found")
	}

	databaseName := os.Getenv(databaseName)
	if len(databaseName) == 0 {
		return nil, errors.New("db DatabaseName not found")
	}

	return &DbConfig{
		Host:         host,
		Port:         port,
		Driver:       driver,
		User:         user,
		Password:     password,
		SslMode:      sslMode,
		DatabaseName: databaseName,
	}, nil
}

func (cfg *DbConfig) ConnectString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DatabaseName,
		cfg.SslMode,
	)
}
