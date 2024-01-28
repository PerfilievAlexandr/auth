package authMapper

import (
	"github.com/PerfilievAlexandr/auth/internal/api/grpc/auth/dtoGrpcAuth"
	proto "github.com/PerfilievAlexandr/auth/pkg/auth_v1"
)

func MapToLoginRequest(req *proto.LoginRequest) authGrpcDto.LoginRequest {
	return authGrpcDto.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}
}
