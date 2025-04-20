package auth

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/obeismo/auth/internal/model"
	"github.com/obeismo/auth/internal/repository"
	"github.com/obeismo/auth/internal/repository/auth/converter"
	modelRepo "github.com/obeismo/auth/internal/repository/auth/model"
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

func (r *repo) Create(ctx context.Context, auth *model.Auth) (int64, error) {
	builder := squirrel.Insert(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(auth.Name, auth.Email, auth.Password, auth.IsAdmin).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var id int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Auth, error) {
	builder := squirrel.Select(idColumn, nameColumn, emailColumn, passwordColumn, roleColumn,
		createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(squirrel.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var auth modelRepo.Auth
	err = r.db.QueryRow(ctx, query, args...).Scan(&auth)
	if err != nil {
		return nil, err
	}

	return converter.ToAuthFromRepo(&auth), nil
}
