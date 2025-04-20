package converter

import (
	"github.com/obeismo/auth/internal/model"
	modelRepo "github.com/obeismo/auth/internal/repository/auth/model"
)

func UserDataFromRepo(auth *modelRepo.UserData) *model.UserData {
	return &model.UserData{
		ID:        auth.ID,
		Name:      auth.Name,
		Email:     auth.Email,
		IsAdmin:   auth.IsAdmin,
		CreatedAt: auth.CreatedAt,
		UpdatedAt: auth.UpdatedAt,
	}
}
