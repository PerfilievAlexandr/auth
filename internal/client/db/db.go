package db

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Client interface {
	ScanOneContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) pgx.Row
	Ping(ctx context.Context) error
	Close()
}
