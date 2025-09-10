package entity

import "github.com/jackc/pgx/v5/pgtype"

type Profile struct {
	Id pgtype.Int4 `db:"id"`
}
