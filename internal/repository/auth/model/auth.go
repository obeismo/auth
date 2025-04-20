package model

import (
	"database/sql"
	"time"
)

type Auth struct {
	ID              int64
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	IsAdmin         bool
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}
