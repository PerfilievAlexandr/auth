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
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return password, err
	}

	return string(hash), nil
}

func (p passwordServiceImpl) CompareWithHashedPassword(_ context.Context, newPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(strings.TrimSpace(newPassword)), []byte(hashedPassword))
	if err != nil {
		return false
	}

	return true
}

func (p passwordServiceImpl) CompareWithConfirmPassword(_ context.Context, password string, confirmPassword string) bool {
	return strings.TrimSpace(password) == strings.TrimSpace(confirmPassword)
}
