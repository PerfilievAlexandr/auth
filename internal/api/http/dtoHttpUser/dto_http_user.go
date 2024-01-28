package dtoHttpUser

import (
	"database/sql"
	"time"
)

type SignUpRequest struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
}

type UpdateRequest struct {
	Name  string
	Email string
}

type UserResponse struct {
	Id        int64
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
