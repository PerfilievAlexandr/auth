package mapper

import (
	"github.com/PerfilievAlexandr/auth/internal/domain"
	dbUser "github.com/PerfilievAlexandr/auth/internal/repository/user/dto"
)

func ToUserFromUserDb(userDb *dbUser.UserDb) *domain.User {
	return &domain.User{
		Id:        userDb.Id,
		Name:      userDb.Name,
		Email:     userDb.Email,
		Role:      userDb.Role,
		CreatedAt: userDb.CreatedAt,
		UpdatedAt: userDb.UpdatedAt,
	}
}
