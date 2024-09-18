package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Makovey/microservice_auth/internal/adapter"
	"github.com/Makovey/microservice_auth/internal/repository/user"
	"github.com/Makovey/microservice_auth/internal/service"
	proto "github.com/Makovey/microservice_auth/pkg/user/v1"
)

const (
	dbDSN    = "host=176.114.66.95 port=5432 dbname=postgres user=admin password=admin sslmode=disable"
	grpcPort = 3000
)

type server struct {
	proto.UnimplementedUserV1Server
	userService service.UserService
}

func (s *server) Create(ctx context.Context, req *proto.User) (*proto.CreateResponse, error) {
	id, err := s.userService.Create(ctx, adapter.ToUserFromProto(req))
	if err != nil {
		return nil, err
	}

	log.Printf("Created user with id %d", id)

	return &proto.CreateResponse{Id: id}, nil
}

func (s *server) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	res, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("Get user with id %d", res.ID)

	return adapter.ToProtoFromUser(res), nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err = pool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}
	defer pool.Close()

	repo := user.NewRepository(pool)
	userSrv := service.UserService(repo)

	s := grpc.NewServer()
	reflection.Register(s)
	proto.RegisterUserV1Server(s, &server{userService: userSrv})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
