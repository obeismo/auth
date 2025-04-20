package main

import (
	"context"
	"github.com/obeismo/auth/internal/repository"
	"log"

	desc "github.com/obeismo/auth/pkg/auth/v1"
)

type server struct {
	desc.UnimplementedAuthV1Server
	authRepository repository.AuthRepository
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.authRepository.Create(ctx, req.GetInfo())
	if err != nil {
		return nil, err
	}

	log.Printf("inserted auth with id: %s", id)

	return &desc.CreateResponse{
		id: id,
	}, nil
}

func main() {
	ctx := context.Background()

	//a, err := app.NewApp(ctx)
	//if err != nil {
	//	log.Fatalf("failed to init app: %s", err.Error())
	//}
	//
	//err = a.Run()
	//if err != nil {
	//	log.Fatalf("failed to run app: %s", err.Error())
	//}
}
