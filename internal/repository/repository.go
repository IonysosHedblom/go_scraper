package repository

import "database/sql"

type repository struct {
	db *sql.DB
}

func NewRepository(dialect, dsn string) (*repository, error) {
	db, err := sql.Open(dialect, dsn)

	if err != nil {
		return nil, err
	}

	return &repository{db}, nil
}

// func (r *repository) Get(query string) int {

// }
