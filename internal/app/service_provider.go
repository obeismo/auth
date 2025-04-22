package app

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/obeismo/auth/internal/api/auth"
	"github.com/obeismo/auth/internal/closer"
	"github.com/obeismo/auth/internal/config"
	"github.com/obeismo/auth/internal/repository"
	authRepository "github.com/obeismo/auth/internal/repository/auth"
	"github.com/obeismo/auth/internal/service"
	authService "github.com/obeismo/auth/internal/service/auth"
	"log"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgPool         *pgxpool.Pool
	authRepository repository.AuthRepository

	authService service.AuthService

	authServer *auth.Server
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failet to get grpc config: %s", err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to db: %s", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping db: %s", err)
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pgPool = pool
	}

	return s.pgPool
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authRepository.NewRepository(s.PgPool(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.AuthRepository(ctx))
	}

	return s.authService
}

func (s *serviceProvider) AuthServer(ctx context.Context) *auth.Server {
	if s.authServer == nil {
		s.authServer = auth.NewServer(s.AuthService(ctx))
	}

	return s.authServer
}
