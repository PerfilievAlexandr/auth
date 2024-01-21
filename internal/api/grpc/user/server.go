package user

import (
	"context"
	"github.com/PerfilievAlexandr/auth/internal/api/grpc/user/mapper"
	"github.com/PerfilievAlexandr/auth/internal/service"
	proto "github.com/PerfilievAlexandr/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	proto.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Server {
	return &Server{
		userService: userService,
	}
}

func (s *Server) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	userId, err := s.userService.Create(ctx, mapper.MapToCreateUser(req))

	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{
		Id: userId,
	}, nil
}

func (s *Server) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	user, err := s.userService.Get(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	return mapper.MapToUserApi(user), nil
}

func (s *Server) Update(ctx context.Context, req *proto.UpdateRequest) (*emptypb.Empty, error) {
	empty, err := s.userService.Update(ctx, mapper.MapToUpdateUser(req))

	if err != nil {
		return nil, err
	}

	return empty, nil
}

func (s *Server) Delete(ctx context.Context, req *proto.DeleteRequest) (*emptypb.Empty, error) {
	empty, err := s.userService.Delete(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return empty, nil
}
