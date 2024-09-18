package user

import (
	"context"

	"github.com/Makovey/microservice_auth/internal/model"
)

func (s *service) Create(ctx context.Context, user *model.User) (int64, error) {
	id, err := s.repository.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}
