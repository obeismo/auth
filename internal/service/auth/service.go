package auth

import (
	"github.com/obeismo/auth/internal/client/db"
	"github.com/obeismo/auth/internal/repository"
	"github.com/obeismo/auth/internal/service"
)

type serv struct {
	authRepository repository.AuthRepository
	txManager      db.TxManager
}

func NewService(authRepository repository.AuthRepository, txManager db.TxManager) service.AuthService {
	return &serv{
		authRepository: authRepository,
		txManager:      txManager,
	}
}
