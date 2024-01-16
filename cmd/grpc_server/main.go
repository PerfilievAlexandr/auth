package main

import (
	userServer "auth/internal/api/grpc/user"
	"auth/internal/config"
	userRepository "auth/internal/repository/user"
	userService "auth/internal/service/user"
	proto "auth/pkg/user_v1"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	ctx := context.Background()

	// Считываем переменные окружения
	err := config.Load(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, cfg.DbConfig.ConnectString())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", cfg.GRPCConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	repository := userRepository.NewRepository(pool)
	service := userService.NewUserService(repository)
	server := userServer.NewImplementation(service)

	s := grpc.NewServer()
	reflection.Register(s)
	proto.RegisterUserV1Server(s, server)

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
