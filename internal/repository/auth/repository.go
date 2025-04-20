package auth

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/obeismo/auth/internal/repository"
)

const (
	tableName = "auth"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.AuthRepository {
	return &repo{db: db}
}

//func (r *repo) Create(ctx context.Context, auth *)
