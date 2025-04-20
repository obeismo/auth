package model

import (
	"time"
)

type NewUser struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            bool
}

type UserData struct {
	ID        int64
	Name      string
	Email     string
	Role      bool
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UpdatedUserData struct {
	Name     string
	Email    string
	Password string
	Role     bool
}
