package service

import (
	"context"

	"github.com/Makovey/microservice_auth/internal/model"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
