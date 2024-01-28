package dtoUserDb

import (
	"database/sql"
	"time"
)

type UserDb struct {
	Id        int64        `db:"id"`
	Name      string       `db:"name"`
	Password  string       `db:"password"`
	Email     string       `db:"email"`
	Role      string       `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
