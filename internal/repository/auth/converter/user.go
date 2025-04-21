package converter

import (
	"github.com/obeismo/auth/internal/constants"
	"github.com/obeismo/auth/internal/model"
	modelRepo "github.com/obeismo/auth/internal/repository/auth/model"
	desc "github.com/obeismo/auth/pkg/auth/v1"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	if user == nil {
		return &model.User{}
	}

	return &model.User{
		ID:        user.ID,
		Info:      *ToUserDataFromRepo(user),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserDataFromRepo(user *modelRepo.User) *model.UserInfo {
	if user == nil {
		return &model.UserInfo{}
	}

	return &model.UserInfo{
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: "",
		Role:            RoleFromString(user.Role),
	}
}

func RoleFromString(s string) desc.Role {
	switch s {
	case constants.USER:
		return desc.Role_USER
	case constants.ADMIN:
		return desc.Role_ADMIN
	default:
		return desc.Role_ROLE_UNSPECIFIED
	}
}
