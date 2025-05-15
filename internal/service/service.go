package service

import (
	"context"

	"github.com/obeismo/auth/internal/model"
)

type AuthService interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
