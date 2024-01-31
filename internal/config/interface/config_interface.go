package confog_interface

import "time"

type GrpcServerConfig interface {
	Address() string
}

type HttpServerConfig interface {
	Address() string
}

type DatabaseConfig interface {
	ConnectString() string
}

type JwtConfig interface {
	RefreshTokenSecret() string
	AccessTokenSecret() string
	RefreshTokenExpirationMinutes() time.Duration
	AccessTokenExpirationMinutes() time.Duration
}

type PrometheusServerConfig interface {
	Address() string
}
