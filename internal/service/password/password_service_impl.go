package passwordService

import (
	"context"
	"github.com/PerfilievAlexandr/auth/internal/service"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type passwordServiceImpl struct{}

func NewPasswordService() service.PasswordService {
	return &passwordServiceImpl{}
}

func (p passwordServiceImpl) HashAndSaltPassword(_ context.Context, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(password)), bcrypt.MinCost)
	if err != nil {
		return password, err
	}

	return string(hash), nil
}

func (p passwordServiceImpl) CompareWithHashedPassword(_ context.Context, dbPassword string, newPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(strings.TrimSpace(newPassword)))
	if err != nil {
		return false
	}

	return true
}

func (p passwordServiceImpl) CompareWithConfirmPassword(_ context.Context, password string, confirmPassword string) bool {
	return strings.TrimSpace(password) == strings.TrimSpace(confirmPassword)
}
