package user

import (
	"github.com/Makovey/microservice_auth/internal/repository"
	def "github.com/Makovey/microservice_auth/internal/service"
)

type service struct {
	repository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) def.UserService {
	return &service{repository: userRepository}
}
