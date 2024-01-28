package userRepository

import (
	"context"
	"fmt"
	"github.com/PerfilievAlexandr/auth/internal/api/http/dtoHttpUser"
	"github.com/PerfilievAlexandr/auth/internal/domain"
	"github.com/PerfilievAlexandr/auth/internal/dto"
	"github.com/PerfilievAlexandr/auth/internal/repository"
	"github.com/PerfilievAlexandr/auth/internal/repository/user/dtoUserDb"
	"github.com/PerfilievAlexandr/auth/internal/repository/user/mapper"
	"github.com/PerfilievAlexandr/platform_common/pkg/db"
	"time"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db}
}

func (s *repo) Create(ctx context.Context, req dto.CreateUser) (int64, error) {
	var userId int64
	query := fmt.Sprintf("INSERT INTO users (name, email, password, role, created_at) values ($1, $2, $3, $4, $5) RETURNING id")
	err := s.db.ScanOneContext(ctx, &userId, query, req.Name, req.Email, req.Password, req.Role, time.Now())
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *repo) Get(ctx context.Context, userId int64) (*domain.User, error) {
	query := fmt.Sprintf(`SELECT s.id, s.name, s.email, s.role, s.created_at, s.updated_at FROM users s WHERE s.id = $1`)
	var dbUser = dtoUserDb.UserDb{}
	err := s.db.ScanOneContext(ctx, &dbUser, query, userId)
	if err != nil {
		return nil, err
	}

	return mapper.ToUserFromUserDb(&dbUser), nil
}

func (s *repo) GetAll(ctx context.Context) ([]*domain.User, error) {
	query := fmt.Sprintf(`SELECT s.id, s.name, s.email, s.role, s.created_at, s.updated_at FROM users s LIMIT 50`)
	var dbUsers []dtoUserDb.UserDb
	err := s.db.ScanAllContext(ctx, &dbUsers, query)
	if err != nil {
		return nil, err
	}

	var domainUsers []*domain.User
	for _, dbUser := range dbUsers {
		mappedUser := mapper.ToUserFromUserDb(&dbUser)
		domainUsers = append(domainUsers, mappedUser)
	}

	return domainUsers, nil
}

func (s *repo) Update(ctx context.Context, userId int64, req dtoHttpUser.UpdateRequest) error {
	query := fmt.Sprintf("UPDATE users SET name=$2, email=$3, updated_at=$4 WHERE id=$1")
	_, err := s.db.ExecContext(ctx, query, userId, req.Name, req.Email, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (s *repo) Delete(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("DELETE FROM users WHERE id=$1")
	_, err := s.db.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	return nil
}
