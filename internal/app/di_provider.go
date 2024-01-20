package app

import (
	"auth/internal/api/grpc/user"
	"auth/internal/client/db"
	"auth/internal/client/db/pg"
	"auth/internal/client/db/transaction"
	"auth/internal/closer"
	"auth/internal/config"
	"auth/internal/repository"
	userRepository "auth/internal/repository/user"
	"auth/internal/service"
	userService "auth/internal/service/user"
	"context"
	"log"
)

type diProvider struct {
	config         *config.Config
	db             db.Client
	txManager      db.TxManager
	userRepository repository.UserRepository
	userService    service.UserService
	server         *user.Server
}

func newProvider() *diProvider {
	return &diProvider{}
}

func (s *diProvider) Config(ctx context.Context) *config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig(ctx)
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *diProvider) DbClient(ctx context.Context) db.Client {
	if s.db == nil {

		dbPool, err := pg.New(ctx, s.Config(ctx).DbConfig.ConnectString())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = dbPool.Ping(ctx)
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		closer.Add(func() error {
			dbPool.Close()
			return nil
		})

		s.db = dbPool
	}

	return s.db
}

func (s *diProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DbClient(ctx))
	}

	return s.txManager
}

func (s *diProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DbClient(ctx))
	}

	return s.userRepository
}

func (s *diProvider) UserService(ctx context.Context) repository.UserRepository {
	if s.userService == nil {
		s.userService = userService.NewUserService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}

func (s *diProvider) UserServer(ctx context.Context) *user.Server {
	if s.server == nil {
		s.server = user.NewImplementation(s.UserService(ctx))
	}

	return s.server
}
