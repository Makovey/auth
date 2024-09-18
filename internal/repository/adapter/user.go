package adapter

import (
	repo "github.com/Makovey/microservice_auth/internal/repository/model"
	proto "github.com/Makovey/microservice_auth/pkg/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProtoFromRepo(user *repo.RepoUser) *proto.GetResponse {
	var updatedAt *timestamppb.Timestamp

	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &proto.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      proto.Role(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
