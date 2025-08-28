package store

import (
	"context"
	"fmt"
	"test/internal/config"
	"test/internal/store/entity"

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

func (d Db) GetProfile(ctx context.Context, lang string) (entity.Profile, error) {
	q := "select * from get_profile($1)"

	var m entity.Profile

	err := d.pool.QueryRow(ctx, q, lang).Scan(
		&m.Id,
		&m.Email,
		&m.Phone,
		&m.Salary,
		&m.SalaryExpectation,
		&m.FullName,
		&m.Residence,
		&m.Skills,
	)

	if err != nil {
		return m, err
	}

	return m, nil
}

func (d Db) GetExperience(ctx context.Context, lang string) (entity.Experience, error) {
	q := "select * from get_experience($1)"

	var m entity.Experience

	err := d.pool.QueryRow(ctx, q, lang).Scan(
		&m.Id,
		&m.CompanyName,
		&m.PositionName,
		&m.Skills,
		&m.StartDt,
		&m.EndDt,
	)

	if err != nil {
		return m, err
	}

	return m, nil
}

func (d Db) GetEducation(ctx context.Context, lang string) (entity.Education, error) {
	q := "select * from get_education($1)"

	var m entity.Education

	err := d.pool.QueryRow(ctx, q, lang).Scan(
		&m.Id,
		&m.UniversityName,
		&m.Achievements,
		&m.StartDt,
		&m.EndDt,
	)

	if err != nil {
		return m, err
	}

	return m, nil
}

func (d Db) GetLanguages(ctx context.Context, lang string) ([]entity.Language, error) {
	q := "select * from get_languages($1)"

	var langs []entity.Language

	rows, err := d.pool.Query(ctx, q, lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var l entity.Language

		if e := rows.Scan(&l.Id, &l.Txt, &l.Level); e != nil {
			return nil, e
		}

		langs = append(langs, l)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return langs, nil
}

func (d Db) GetProjects(ctx context.Context, lang string) ([]entity.Project, error) {
	q := "select * from get_projects($1)"

	var projects []entity.Project

	rows, err := d.pool.Query(ctx, q, lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p entity.Project

		if e := rows.Scan(&p.Id, &p.ProjectName, &p.Link, &p.LinkTxt); e != nil {
			return nil, e
		}

		projects = append(projects, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
