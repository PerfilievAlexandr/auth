package app

import (
	"context"
	"github.com/PerfilievAlexandr/auth/internal/api/grpc/interceptor"
	"github.com/PerfilievAlexandr/auth/internal/config"
	rateLimiter "github.com/PerfilievAlexandr/auth/internal/helpers/rate_limiter"
	"github.com/PerfilievAlexandr/auth/internal/logger"
	"github.com/PerfilievAlexandr/auth/internal/metrics"
	"github.com/PerfilievAlexandr/auth/internal/tracing"
	protoAccess "github.com/PerfilievAlexandr/auth/pkg/access_v1"
	protoAuth "github.com/PerfilievAlexandr/auth/pkg/auth_v1"
	"github.com/PerfilievAlexandr/platform_common/pkg/closer"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sony/gobreaker"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

type App struct {
	diProvider *diProvider
	grpcServer *grpc.Server
	httpServer *http.Server
	prometheus *http.Server
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
			logger.Fatal("failed to run GRPC grpcAuthServer", zap.Any("err", err))
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runHttpServer(ctx)
		if err != nil {
			logger.Fatal("failed to run HTTP httpAuthServer", zap.Any("err", err))
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runPrometheus(ctx)
		if err != nil {
			logger.Fatal("failed to run prometheusServer", zap.Any("err", err))
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initLogger,
		a.initTrace,
		a.initConfig,
		a.initProvider,
		a.initGrpcServer,
		a.initHttpServer,
		a.initPrometheus,
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
	limiter := rateLimiter.NewTokenBucketLimiter(ctx, 100, time.Second)
	circuitBreaker := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "auth-service",
		MaxRequests: 3,
		Timeout:     5 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return failureRatio >= 0.5
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			logger.Info("Circuit Breaker: %s, changed from %v, to %v\n", zap.String("name", name), zap.Any("changed from", from), zap.Any("changed to", to))
		},
	})

	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interceptor.NewRateLimiterInterceptor(limiter).Unary,
				interceptor.NewCircuitBreakerInterceptor(circuitBreaker).Unary,
				interceptor.ValidateInterceptor,
				interceptor.LogInterceptor,
				interceptor.MetricsInterceptor,
				otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
			),
		),
	)

	reflection.Register(a.grpcServer)
	protoAuth.RegisterAuthV1Server(a.grpcServer, a.diProvider.GrpcAuthServer(ctx))
	protoAccess.RegisterAccessV1Server(a.grpcServer, a.diProvider.GrpcAccessServer(ctx))

	return nil
}

func (a *App) runGrpcServer(ctx context.Context) error {
	logger.Info("GRPC server is running on:", zap.String("host:port", a.diProvider.config.GRPCConfig.Address()))

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
	logger.Info("HTTP server is running on:", zap.String("host:port", a.diProvider.config.HttpConfig.Address()))

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	stdout := zapcore.AddSync(os.Stdout)
	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	core := zapcore.NewCore(consoleEncoder, stdout, zap.InfoLevel)
	logger.Init(core)

	return nil
}

func (a *App) initPrometheus(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	a.prometheus = &http.Server{
		Addr:    a.diProvider.config.PrometheusConfig.Address(),
		Handler: mux,
	}

	return metrics.Init(ctx)
}

func (a *App) runPrometheus(_ context.Context) error {
	logger.Info("Prometheus server is running on:", zap.String("host:port", a.diProvider.config.PrometheusConfig.Address()))

	err := a.prometheus.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initTrace(_ context.Context) error {
	tracing.Init(logger.Logger(), "auth")

	return nil
}
