package auth

import (
	"context"
	authMapper "github.com/PerfilievAlexandr/auth/internal/api/grpc/auth/mapper"
	"github.com/PerfilievAlexandr/auth/internal/service"
	proto "github.com/PerfilievAlexandr/auth/pkg/auth_v1"
)

type Server struct {
	proto.UnimplementedAuthV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Server {
	return &Server{
		authService: authService,
	}
}

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	refreshToken, err := s.authService.Login(ctx, authMapper.MapToLoginRequest(req))

	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		RefreshToken: refreshToken,
	}, nil
}

func (s *Server) GetRefreshToken(ctx context.Context, req *proto.GetRefreshTokenRequest) (*proto.GetRefreshTokenResponse, error) {
	refreshToken, err := s.authService.GetRefreshToken(ctx, req.RefreshToken)

	if err != nil {
		return nil, err
	}

	return &proto.GetRefreshTokenResponse{
		RefreshToken: refreshToken,
	}, nil
}

func (s *Server) GetAccessToken(ctx context.Context, req *proto.GetAccessTokenRequest) (*proto.GetAccessTokenResponse, error) {
	accessToken, err := s.authService.GetAccessToken(ctx, req.RefreshToken)

	if err != nil {
		return nil, err
	}

	return &proto.GetAccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}
