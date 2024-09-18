package user

import (
	"context"
	"log"

	"github.com/Makovey/microservice_auth/internal/adapter"
	proto "github.com/Makovey/microservice_auth/pkg/user/v1"
)

func (s *Server) Create(ctx context.Context, req *proto.User) (*proto.CreateResponse, error) {
	id, err := s.userService.Create(ctx, adapter.ToUserFromProto(req))
	if err != nil {
		return nil, err
	}

	log.Printf("Created user with id %d", id)

	return &proto.CreateResponse{Id: id}, nil
}
