package accessService

import (
	"context"
	"errors"
	"github.com/PerfilievAlexandr/auth/internal/service"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	authPrefix = "Bearer "
)

type accessService struct {
	jwtService service.JwtService
}

func NewAccessService(
	jwtService service.JwtService,
) service.AccessService {
	return &accessService{
		jwtService: jwtService,
	}
}

func (a *accessService) Check(ctx context.Context, _ string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return errors.New("invalid authorization header format")
	}
	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)
	_, err := a.jwtService.VerifyAccessToken(ctx, accessToken)
	if err != nil {
		return errors.New("access token is invalid")
	}

	// can check access to endpoint here for example get map from db with key - endpoint and value - role
	// and compare value with claims role

	return nil
}
