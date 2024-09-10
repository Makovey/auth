package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	m "github.com/Makovey/microservice_auth/pkg/user/v1"
)

const grpcPort = 3000

type server struct {
	m.UnimplementedUserV1Server
}

func (s *server) Create(ctx context.Context, in *m.User) (*m.CreateResponse, error) {
	return &m.CreateResponse{Id: 1}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	m.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
