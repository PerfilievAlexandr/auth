package repository

import (
	"context"
	"github.com/PerfilievAlexandr/auth/internal/api/http/dtoHttpUser"
	"github.com/PerfilievAlexandr/auth/internal/domain"
	"github.com/PerfilievAlexandr/auth/internal/dto"
)

type UserRepository interface {
	Create(ctx context.Context, req dto.CreateUser) (int64, error)
	GetById(ctx context.Context, userId int64) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	GetAll(ctx context.Context) ([]*domain.User, error)
	Update(ctx context.Context, userId int64, req dtoHttpUser.UpdateRequest) error
	Delete(ctx context.Context, id int64) error
}
