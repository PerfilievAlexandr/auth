package accessService

import (
	"context"
	"errors"
	"github.com/PerfilievAlexandr/auth/internal/dto"
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

func (a *accessService) Check(ctx context.Context) (*dto.JwtClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return nil, errors.New("invalid authorization header format")
	}
	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)
	claims, err := a.jwtService.VerifyAccessToken(ctx, accessToken)
	if err != nil {
		return nil, errors.New("access token is invalid")
	}

	return claims, nil
}
