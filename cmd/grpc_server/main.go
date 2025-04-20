package main

import (
	"context"
	"github.com/obeismo/auth/internal/repository"

	desc "github.com/obeismo/auth/pkg/auth/v1"
)

type server struct {
	desc.UnimplementedAuthV1Server
	authRepository repository.AuthRepository
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	return nil, nil
}

func main() {
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
