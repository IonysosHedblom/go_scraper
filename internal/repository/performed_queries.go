package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
)

type performedQueriesStore struct {
	db *sql.DB
}

func NewPerformedQueriesStore(db *sql.DB) *performedQueriesStore {
	return &performedQueriesStore{
		db: db,
	}
}

func (r *performedQueriesStore) GetByQuery(query string) (*entity.PerformedQuery, error) {
	pq := new(entity.PerformedQuery)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, "SELECT query_id, query FROM performed_queries WHERE query = $1", query).Scan(&pq.Id, &pq.Query)

	if err != nil {
		return nil, err
	}

	return pq, nil
}

func (r *performedQueriesStore) GetById(id int) (*entity.PerformedQuery, error) {
	pq := new(entity.PerformedQuery)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, "SELECT query_id, query FROM performed_queries WHERE query_id = $1", id).Scan(&pq.Id, &pq.Query)

	if err != nil {
		return nil, err
	}

	return pq, nil
}

func (r *performedQueriesStore) Create(query string) (*int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbquery := "INSERT INTO performed_queries (query) VALUES ($1)"

	statement, err := r.db.PrepareContext(ctx, dbquery)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	res, err := statement.ExecContext(ctx, query)

	if err != nil {
		return nil, err
	}

	queryId, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &queryId, nil
}
