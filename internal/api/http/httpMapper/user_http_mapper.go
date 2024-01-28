package httpMapper

import (
	"github.com/PerfilievAlexandr/auth/internal/api/http/dtoHttpUser"
	"github.com/PerfilievAlexandr/auth/internal/domain"
)

func MapUserToApiDto(user *domain.User) *dtoHttpUser.UserResponse {
	return &dtoHttpUser.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
