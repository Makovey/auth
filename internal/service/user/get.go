package user

import (
	"context"

	"github.com/Makovey/microservice_auth/internal/model"
)

func (s *service) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
