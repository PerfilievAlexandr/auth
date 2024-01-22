package dto

import (
	"database/sql"
	"time"
)

type UserApi struct {
	Id        int64
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
