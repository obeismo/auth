package auth

import (
	"context"
	"github.com/obeismo/auth/internal/model"
	"log"
)

func (s *Server) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	id, err := s.authService.Create(ctx, info)
	if err != nil {
		return 0, err
	}

	log.Printf("inserted user with id: %d", id)

	return id, nil
}
