package app

import (
	"context"
	"github.com/PerfilievAlexandr/auth/internal/config"
	proto "github.com/PerfilievAlexandr/auth/pkg/user_v1"
	"github.com/PerfilievAlexandr/platform_common/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"sync"
)

type App struct {
	diProvider *diProvider
	grpcServer *grpc.Server
	httpServer *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()

		err := a.runGrpcServer(ctx)
		if err != nil {
			log.Fatalf("failed to run GRPC grpcServer: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runHttpServer(ctx)
		if err != nil {
			log.Fatalf("failed to run HTTP grpcServer: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initProvider,
		a.initGrpcServer,
		a.initHttpServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initProvider(_ context.Context) error {
	a.diProvider = newProvider()
	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	proto.RegisterUserV1Server(a.grpcServer, a.diProvider.GrpcServer(ctx))

	return nil
}

func (a *App) runGrpcServer(ctx context.Context) error {
	log.Printf("GRPC server is running on %s", a.diProvider.Config(ctx).GRPCConfig.Address())

	list, err := net.Listen("tcp", a.diProvider.Config(ctx).GRPCConfig.Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initHttpServer(ctx context.Context) error {
	a.diProvider.HttpHandler(ctx)
	a.httpServer = &http.Server{
		Addr:    a.diProvider.config.HttpConfig.Address(),
		Handler: a.diProvider.httpHandler.InitRoutes(),
	}

	return nil
}

func (a *App) runHttpServer(_ context.Context) error {
	log.Printf("HTTP server is running on %s", a.diProvider.config.HttpConfig.Address())

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
