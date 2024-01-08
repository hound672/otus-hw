package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(config *Config) (*pgxpool.Pool, func(), error) {
	ctx := context.Background()

	uri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.Username, config.Password, config.Host, config.Port, config.Database,
	)

	pool, err := pgxpool.New(ctx, uri)
	if err != nil {
		return nil, nil, fmt.Errorf("pgxpool.New: %w", err)
	}

	cleanup := func() {
		pool.Close()
	}

	return pool, cleanup, nil
}
