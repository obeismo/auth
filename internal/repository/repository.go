package repository

import (
	"context"
	"github.com/obeismo/auth/internal/model"
)

type AuthRepository interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
