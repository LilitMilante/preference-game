package bootstrap

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresClient(cfg *Config) (*pgxpool.Pool, error) {
	ctx := context.Background()

	conn, err := pgxpool.New(ctx, cfg.PostgresURL)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
