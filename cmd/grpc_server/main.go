package main

import (
	"auth/internal/config"
	desc "auth/pkg/user_v1"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"
)

type server struct {
	desc.UnimplementedUserV1Server
	pool *pgxpool.Pool
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	var id int64
	var name, email, role string
	var createdAt time.Time
	var updatedAt sql.NullTime
	query := fmt.Sprintf(`SELECT s.id, s.name, s.email, s.role, s.created_at, s.updated_at FROM users s WHERE s.id = $1`)
	row := s.pool.QueryRow(ctx, query, req.GetId())
	err := row.Scan(&id, &name, &email, &role, &createdAt, &updatedAt)
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}

	var updatedAtTime *timestamppb.Timestamp
	if updatedAt.Valid {
		updatedAtTime = timestamppb.New(updatedAt.Time)
	}

	return &desc.GetResponse{
		Id:        id,
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: updatedAtTime,
	}, nil
}

func (s *server) Create(ctx context.Context, req *desc.CreteRequest) (*desc.CreateResponse, error) {
	var userId int64
	createUserQuery := fmt.Sprintf("INSERT INTO users (name, email, password, role, created_at) values ($1, $2, $3, $4, $5) RETURNING id")
	err := s.pool.QueryRow(ctx, createUserQuery, req.Name, req.Email, req.Password, req.Role, time.Now()).Scan(&userId)
	if err != nil {
		log.Fatalf("failed to insert user: %v", err)
	}

	return &desc.CreateResponse{Id: userId}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	updateUserQuery := fmt.Sprintf("UPDATE users SET name=$2, email=$3, updated_at=$4 WHERE id=$1")
	_, err := s.pool.Exec(ctx, updateUserQuery, req.Id, req.Name.Value, req.Email.Value, time.Now())
	if err != nil {
		log.Fatalf("failed to update user: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	deleteUserQuery := fmt.Sprintf("DELETE FROM users WHERE id=$1")
	_, err := s.pool.Exec(ctx, deleteUserQuery, req.Id)
	if err != nil {
		log.Fatalf("failed to delete user: %v", err)
	}

	return &emptypb.Empty{}, nil
}

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

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{pool: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
