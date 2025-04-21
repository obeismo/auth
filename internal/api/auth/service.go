package auth

import (
	"github.com/obeismo/auth/internal/service"
	desc "github.com/obeismo/auth/pkg/auth/v1"
)

type Server struct {
	desc.UnimplementedAuthV1Server
	authService service.AuthService
}

func NewServer(authService service.AuthService) *Server {
	return &Server{
		authService: authService,
	}
}
