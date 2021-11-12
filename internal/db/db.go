package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB(dialect, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dialect, dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
