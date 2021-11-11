package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type repository struct {
	db *sql.DB
}

func NewRepository(dialect, dsn string) (*repository, error) {
	db, err := sql.Open(dialect, dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return &repository{db}, nil
}

func (r *repository) Close() {
	r.db.Close()
}
