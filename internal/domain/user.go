package domain

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int64
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
