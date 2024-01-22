package mapper

import (
	"github.com/PerfilievAlexandr/auth/internal/api/http/dto"
	"github.com/PerfilievAlexandr/auth/internal/domain"
)

func MapUserToApiDto(user *domain.User) *dto.UserApi {
	return &dto.UserApi{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
