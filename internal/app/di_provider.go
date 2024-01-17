package app

import (
	"auth/internal/api/grpc/user"
	"auth/internal/config"
	"auth/internal/repository"
	userRepository "auth/internal/repository/user"
	"auth/internal/service"
	userService "auth/internal/service/user"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type diProvider struct {
	config         *config.Config
	db             *pgxpool.Pool
	userRepository repository.UserRepository
	userService    service.UserService
	server         *user.Server
}

func newProvider() *diProvider {
	return &diProvider{}
}

func (s *diProvider) Config() *config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *diProvider) PgxPool(ctx context.Context) *pgxpool.Pool {
	if s.db == nil {
		pool, err := pgxpool.Connect(ctx, s.config.DbConfig.ConnectString())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		s.db = pool
	}

	return s.db
}

func (s *diProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.PgxPool(ctx))
	}

	return s.userRepository
}

func (s *diProvider) UserService(ctx context.Context) repository.UserRepository {
	if s.userService == nil {
		s.userService = userService.NewUserService(s.UserRepository(ctx))
	}

	return s.userService
}

func (s *diProvider) UserServer(ctx context.Context) *user.Server {
	if s.server == nil {
		s.server = user.NewImplementation(s.UserService(ctx))
	}

	return s.server
}
