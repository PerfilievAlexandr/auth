package mapper

import (
	"github.com/PerfilievAlexandr/auth/internal/domain"
	"github.com/PerfilievAlexandr/auth/internal/repository/user/dtoUserDb"
)

func ToUserFromUserDb(userDb *dtoUserDb.UserDb) *domain.User {
	return &domain.User{
		Id:        userDb.Id,
		Name:      userDb.Name,
		Email:     userDb.Email,
		Role:      userDb.Role,
		CreatedAt: userDb.CreatedAt,
		UpdatedAt: userDb.UpdatedAt,
	}
}
