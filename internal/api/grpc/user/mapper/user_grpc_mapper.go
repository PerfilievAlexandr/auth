package mapper

import (
	"auth/internal/api/grpc/user/dto"
	"auth/internal/domain"
	proto "auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapToUserApi(user *domain.User) *proto.GetResponse {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &proto.GetResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func MapToCreateUser(req *proto.CreteRequest) *dto.CreateRequest {
	return &dto.CreateRequest{
		Name:            req.Name,
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
		Role:            req.Role,
	}
}

func MapToUpdateUser(req *proto.UpdateRequest) *dto.UpdateRequest {
	return &dto.UpdateRequest{
		Id:    req.Id,
		Name:  req.Name.Value,
		Email: req.Email.Value,
	}
}
