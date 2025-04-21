package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	authAPI "github.com/obeismo/auth/internal/api/auth"
	"github.com/obeismo/auth/internal/config"
	authRepository "github.com/obeismo/auth/internal/repository/auth"
	authService "github.com/obeismo/auth/internal/service/auth"
	desc "github.com/obeismo/auth/pkg/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	ctx := context.Background()

	err := config.Load(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failet to load db config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	authRepo := authRepository.NewRepository(pool)
	authSrv := authService.NewService(authRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, authAPI.NewServer(authSrv))

	log.Printf("server listening at: %s", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
