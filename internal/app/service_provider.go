package app

import (
	"context"
	"log"

	"github.com/Makovey/microservice_auth/internal/api/user"
	"github.com/Makovey/microservice_auth/internal/client/db"
	"github.com/Makovey/microservice_auth/internal/client/db/pg"
	"github.com/Makovey/microservice_auth/internal/closer"
	"github.com/Makovey/microservice_auth/internal/config"
	"github.com/Makovey/microservice_auth/internal/repository"
	userRepo "github.com/Makovey/microservice_auth/internal/repository/user"
	"github.com/Makovey/microservice_auth/internal/service"
	userService "github.com/Makovey/microservice_auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbcClient db.Client
	userRepo  repository.UserRepository

	userSrv service.UserService

	userServer *user.Server
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbcClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		if err = cl.DB().Ping(ctx); err != nil {
			log.Fatalf("failed to ping database: %v", err)
		}

		closer.Add(cl.Close)

		s.dbcClient = cl
	}

	return s.dbcClient
}

func (s *serviceProvider) UserRepo(ctx context.Context) repository.UserRepository {
	if s.userRepo == nil {
		s.userRepo = userRepo.NewRepository(s.DBClient(ctx))
	}

	return s.userRepo
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userSrv == nil {
		s.userSrv = userService.NewService(s.UserRepo(ctx))
	}

	return s.userSrv
}

func (s *serviceProvider) UserServer(ctx context.Context) *user.Server {
	if s.userServer == nil {
		s.userServer = user.NewServer(s.UserRepo(ctx))
	}

	return s.userServer
}
