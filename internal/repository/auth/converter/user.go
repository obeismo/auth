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
		UpdatedAt: user.UpdatedAt.Time,
	}
}

func ToUserDataFromRepo(info *modelRepo.User) *model.UserInfo {
	if info == nil {
		return &model.UserInfo{}
	}

	return &model.UserInfo{
		Name:            info.Name,
		Email:           info.Email,
		Password:        info.Password,
		PasswordConfirm: "",
		Role:            RoleFromString(info.Role),
	}
}

func RoleFromString(s string) desc.Role {
	switch s {
	case constants.USER:
		return desc.Role_USER
	case constants.ADMIN:
		return desc.Role_ADMIN
	default:
		return desc.Role_UNKNOWN
	}
}
