package model

import (
	"time"
)

type NewUser struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=5"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
	Role            uint8  `json:"role" validate:"required"`
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
