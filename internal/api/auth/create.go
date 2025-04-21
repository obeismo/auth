package auth

import (
	"context"
	"github.com/obeismo/auth/internal/converter"
	desc "github.com/obeismo/auth/pkg/auth/v1"
	"log"
)

func (s *Server) Create(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	id, err := s.authService.Create(ctx, converter.ToUserInfoServiceFromDesc(req.Info))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with id: %d", id)

	return &desc.CreateUserResponse{
		Id: id,
	}, nil
}
