package repository

import (
	"context"
	"github.com/obeismo/auth/internal/repository/auth/model"
)

type AuthRepository interface {
	Create(ctx context.Context, auth *model.Auth) (int64, error)
	Get(ctx context.Context, id int64) (*model.Auth, error)
}
