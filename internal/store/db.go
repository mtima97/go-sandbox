package store

import (
	"context"
	"fmt"
	"test/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	pool *pgxpool.Pool
}

func NewDbConn(ctx context.Context, cfg config.Cfg) (Db, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return Db{}, err
	}

	if err := pool.Ping(ctx); err != nil {
		return Db{}, err
	}

	return Db{pool: pool}, nil
}

func (d Db) Close() {
	d.pool.Close()
}
