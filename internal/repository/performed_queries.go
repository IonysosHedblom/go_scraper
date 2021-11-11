package repository

import (
	"context"
	"time"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
)

func (r *repository) GetPerformedQueryByQuery(query string) (*entity.PerformedQuery, error) {
	pq := new(entity.PerformedQuery)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, "SELECT query_id, query FROM performed_queries WHERE query = $1", query).Scan(&pq.Id, &pq.Query)

	if err != nil {
		return nil, err
	}

	return pq, nil
}

func (r *repository) GetPerformedQueryById(id int) (*entity.PerformedQuery, error) {
	pq := new(entity.PerformedQuery)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, "SELECT query_id, query FROM performed_queries WHERE query_id = $1", id).Scan(&pq.Id, &pq.Query)

	if err != nil {
		return nil, err
	}

	return pq, nil
}

func (r *repository) CreatePerformedQuery(query string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbquery := "INSERT INTO performed_queries (query) VALUES ($1)"

	statement, err := r.db.PrepareContext(ctx, dbquery)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.ExecContext(ctx, query)
	return err
}
