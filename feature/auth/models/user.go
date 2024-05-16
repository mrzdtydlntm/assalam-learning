package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64        `db:"id"`
	Fullname  string       `db:"fullname"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
