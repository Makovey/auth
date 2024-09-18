package user

import (
	"github.com/Makovey/microservice_auth/internal/service"
	proto "github.com/Makovey/microservice_auth/pkg/user/v1"
)

type Server struct {
	proto.UnimplementedUserV1Server
	userService service.UserService
}

func NewServer(userService service.UserService) *Server {
	return &Server{
		userService: userService,
	}
}
