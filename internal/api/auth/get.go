package auth

import (
	"context"
	"github.com/obeismo/auth/internal/model"
	"log"
)

func (s *Server) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.authService.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	log.Printf("get user with id: %d", user.ID)

	return user, nil
}
