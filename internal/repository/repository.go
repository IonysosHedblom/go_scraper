package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
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

func (r *repository) GetByQuery(query string) (*entity.PerformedQuery, error) {
	pq := new(entity.PerformedQuery)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, "SELECT id, query FROM performed_queries WHERE query = ?", query).Scan(&pq.Id, &pq.Query)

	if err != nil {
		return nil, err
	}

	return pq, nil
}

func (r *repository) Create(pq *entity.PerformedQuery) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO performed_queries(id, query) VALUES (?, ?)"

	statement, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.ExecContext(ctx, pq.Id, pq.Query)
	return err
}
