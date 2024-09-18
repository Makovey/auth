package user

import (
	"context"
	"log"

	"github.com/Makovey/microservice_auth/internal/adapter"
	proto "github.com/Makovey/microservice_auth/pkg/user/v1"
)

func (s *Server) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	res, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("Get user with id %d", res.ID)

	return adapter.ToProtoFromUser(res), nil
}
