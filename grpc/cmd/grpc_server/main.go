package main

import (
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit/v6"
	desc "github.com/obeismo/auth/grpc/pkg/auth/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedAuthV1Server
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	return &desc.GetResponse{
		Info: &desc.AuthResponse{
			Id:        req.GetId(),
			Name:      gofakeit.Name(),
			Email:     gofakeit.Email(),
			Role:      desc.Role_USER,
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: #{err}")
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, &server{})

	log.Printf("server listening at #{lis.Addr()}")

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: #{err}")
	}
}
