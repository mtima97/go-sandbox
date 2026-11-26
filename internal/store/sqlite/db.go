package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Db struct {
	db *sql.DB
}

//goland:noinspection GoResourceLeak
func NewDB(sourceName string) (Db, error) {
	db, err := sql.Open("sqlite", sourceName)
	if err != nil {
		return Db{}, err
	}

	if err = db.Ping(); err != nil {
		_ = db.Close()
		return Db{}, err
	}

	return Db{db: db}, nil
}

func (d Db) Close() error {
	return d.db.Close()
}
