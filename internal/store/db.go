package store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"test/internal/config"
	"time"

	"github.com/jackc/pgx/v5"
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

	_, err = pool.Exec(ctx, "set search_path to v2;")
	if err != nil {
		return Db{}, err
	}

	return Db{pool: pool}, nil
}

func (d Db) Close() {
	d.pool.Close()
}

func (d Db) GetEntityByName(ctx context.Context, name, language string) (json.RawMessage, error) {
	var (
		q string
		r json.RawMessage
	)

	switch name {
	case "profile":
		q = "select * from get_profile($1::text)"
	case "projects":
		q = "select * from get_projects($1::text)"
	case "languages":
		q = "select * from get_languages($1::text)"
	case "experience":
		q = "select * from get_experience($1::text)"
	case "education":
		q = "select * from get_education($1::text)"
	default:
		return nil, errors.New("unknown entity name")
	}

	qctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if err := d.pool.QueryRow(qctx, q, language).Scan(&r); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("entity not found")
		}

		return nil, err
	}

	return r, nil
}
