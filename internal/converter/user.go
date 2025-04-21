package converter

import (
	"github.com/obeismo/auth/internal/constants"
	serviceModel "github.com/obeismo/auth/internal/model"
	desc "github.com/obeismo/auth/pkg/auth/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ToUserDescFromService - ковертер, который преобразует модель сервисного слоя в модель апи (протобаф) слоя
func ToUserDescFromService(user *serviceModel.User) *desc.User {
	if user == nil {
		return nil
	}

	return &desc.User{
		Id:        user.ID,
		Info:      ToUserInfoDescFromService(&user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

// ToUserInfoDescFromService - ковертер, который преобразует модель сервисного слоя в модель апи (протобаф) слоя
func ToUserInfoDescFromService(info *serviceModel.UserInfo) *desc.UserInfo {
	if info == nil {
		return nil
	}

	return &desc.UserInfo{
		Name:            info.Name,
		Email:           info.Email,
		Password:        info.Password,
		PasswordConfirm: info.PasswordConfirm,
		Role:            info.Role,
	}
}

func ToUserInfoServiceFromDesc(info *desc.UserInfo) *serviceModel.UserInfo {
	if info == nil {
		return &serviceModel.UserInfo{}
	}

	return &serviceModel.UserInfo{
		Name:            info.Name,
		Email:           info.Email,
		Password:        info.Password,
		PasswordConfirm: info.PasswordConfirm,
		Role:            info.Role,
	}
}

// ToUpdateUserInfoServiceFromDesc - ковертер, который преобразует модель апи (протобаф) слоя в модель сервисного слоя
func ToUpdateUserInfoServiceFromDesc(info *desc.UpdateUserInfo) *serviceModel.UpdateUserInfo {
	if info == nil {
		return &serviceModel.UpdateUserInfo{}
	}

	return &serviceModel.UpdateUserInfo{
		UserID:          info.Id,
		Name:            checkEmptyOrNil(info.Name),
		OldPassword:     checkEmptyOrNil(info.OldPassword),
		Password:        checkEmptyOrNil(info.Password),
		PasswordConfirm: checkEmptyOrNil(info.PasswordConfirm),
		Role:            &info.Role,
	}
}

// checkEmptyOrNil - функция, которая преобразует *wrapperspb.StringValue в *string
func checkEmptyOrNil(s *wrapperspb.StringValue) *string {
	if s == nil {
		return nil
	}

	str := s.Value
	return &str
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
