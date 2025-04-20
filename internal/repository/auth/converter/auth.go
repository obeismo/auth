package converter

import (
	"github.com/obeismo/auth/internal/repository/auth/model"
)

func ToAuthFromRepo(auth *model.Auth) *model.Auth {
	return &model.Auth{
		ID:              auth.ID,
		Name:            auth.Name,
		Email:           auth.Email,
		Password:        auth.Password,
		PasswordConfirm: auth.PasswordConfirm,
		IsAdmin:         auth.IsAdmin,
		CreatedAt:       auth.CreatedAt,
		UpdatedAt:       auth.UpdatedAt,
	}
}
