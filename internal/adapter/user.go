package adapter

import (
	"github.com/Makovey/microservice_auth/internal/model"
	repo "github.com/Makovey/microservice_auth/internal/repository/model"
	proto "github.com/Makovey/microservice_auth/pkg/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromRepo(user *repo.RepoUser) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      model.Role(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserFromProto(user *proto.User) *model.User {
	return &model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     model.Role(user.Role),
	}
}

func ToProtoFromUser(user *model.User) *proto.GetResponse {
	var updatedAt *timestamppb.Timestamp

	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &proto.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
