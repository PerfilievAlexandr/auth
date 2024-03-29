package access

import (
	"context"
	"github.com/PerfilievAlexandr/auth/internal/service"
	proto "github.com/PerfilievAlexandr/auth/pkg/access_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	proto.UnimplementedAccessV1Server
	accessService service.AccessService
}

func NewImplementation(accessService service.AccessService) *Server {
	return &Server{
		accessService: accessService,
	}
}

func (s *Server) Check(ctx context.Context, _ *emptypb.Empty) (*proto.ClaimsResponse, error) {
	claims, err := s.accessService.Check(ctx)

	if err != nil {
		return nil, err
	}

	return &proto.ClaimsResponse{Username: claims.Username, Role: claims.Role}, nil
}
