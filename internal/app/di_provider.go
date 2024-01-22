package app

import (
	"context"
	"github.com/PerfilievAlexandr/auth/internal/api/grpc/user"
	"github.com/PerfilievAlexandr/auth/internal/api/http"
	"github.com/PerfilievAlexandr/auth/internal/config"
	"github.com/PerfilievAlexandr/auth/internal/repository"
	userRepository "github.com/PerfilievAlexandr/auth/internal/repository/user"
	"github.com/PerfilievAlexandr/auth/internal/service"
	userService "github.com/PerfilievAlexandr/auth/internal/service/user"
	"github.com/PerfilievAlexandr/platform_common/pkg/closer"
	"github.com/PerfilievAlexandr/platform_common/pkg/db"
	"github.com/PerfilievAlexandr/platform_common/pkg/db/pg"
	"github.com/PerfilievAlexandr/platform_common/pkg/db/transaction"
	"log"
)

type diProvider struct {
	config         *config.Config
	db             db.Client
	txManager      db.TxManager
	userRepository repository.UserRepository
	userService    service.UserService
	grpcServer     *user.Server
	httpHandler    *http.Handler
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

func (s *diProvider) GrpcServer(ctx context.Context) *user.Server {
	if s.grpcServer == nil {
		s.grpcServer = user.NewImplementation(s.UserService(ctx))
	}

	return s.grpcServer
}

func (s *diProvider) HttpHandler(ctx context.Context) *http.Handler {
	if s.httpHandler == nil {
		s.httpHandler = http.NewHandler(s.UserService(ctx))
	}

	return s.httpHandler
}
