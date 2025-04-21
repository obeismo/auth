package auth

import (
	"github.com/obeismo/auth/internal/repository"
	"github.com/obeismo/auth/internal/service"
)

type serv struct {
	authRepository repository.AuthRepository
}

func NewService(authRepository repository.AuthRepository) service.AuthService {
	return &serv{
		authRepository: authRepository,
	}
}
