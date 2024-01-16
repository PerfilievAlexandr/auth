package user

import (
	apiDto "auth/internal/api/grpc/user/dto"
	"auth/internal/domain"
	"auth/internal/repository"
	"auth/internal/repository/user/dto"
	"auth/internal/repository/user/mapper"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (s *repo) Create(ctx context.Context, req *apiDto.CreateRequest) (int64, error) {
	var userId int64
	createUserQuery := fmt.Sprintf("INSERT INTO users (name, email, password, role, created_at) values ($1, $2, $3, $4, $5) RETURNING id")
	err := s.db.QueryRow(ctx, createUserQuery, req.Name, req.Email, req.Password, req.Role, time.Now()).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *repo) Get(ctx context.Context, userId int64) (*domain.User, error) {
	query := fmt.Sprintf(`SELECT s.id, s.name, s.email, s.role, s.created_at, s.updated_at FROM users s WHERE s.id = $1`)
	row := s.db.QueryRow(ctx, query, userId)

	var dbUser = dto.UserDb{}
	err := row.Scan(&dbUser.Id, &dbUser.Name, &dbUser.Email, &dbUser.Role, &dbUser.CreatedAt, &dbUser.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return mapper.ToUserFromUserDb(&dbUser), nil
}

func (s *repo) Update(ctx context.Context, req *apiDto.UpdateRequest) (*emptypb.Empty, error) {
	updateUserQuery := fmt.Sprintf("UPDATE users SET name=$2, email=$3, updated_at=$4 WHERE id=$1")
	_, err := s.db.Exec(ctx, updateUserQuery, req.Id, req.Name, req.Email, time.Now())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *repo) Delete(ctx context.Context, userId int64) (*emptypb.Empty, error) {
	deleteUserQuery := fmt.Sprintf("DELETE FROM users WHERE id=$1")
	_, err := s.db.Exec(ctx, deleteUserQuery, userId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
