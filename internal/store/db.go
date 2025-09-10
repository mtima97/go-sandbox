package store

import (
	"context"
	"encoding/json"
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

func (d Db) GetProfile(ctx context.Context, lang string) (json.RawMessage, error) {
	q := "select * from get_profile($1)"

	var raw json.RawMessage

	if err := d.pool.QueryRow(ctx, q, lang).Scan(&raw); err != nil {
		return nil, err
	}

	return raw, nil
}
