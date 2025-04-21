package auth

import (
	"context"
	"github.com/obeismo/auth/internal/converter"
	desc "github.com/obeismo/auth/pkg/auth/v1"
	"log"
)

func (s *Server) Get(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	user, err := s.authService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("get user with id: %d", user.ID)

	return &desc.GetUserResponse{
		User: converter.ToUserDescFromService(user),
	}, nil
}
