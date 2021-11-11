package repository

import (
	"database/sql"

	"github.com/ionysoshedblom/go_scraper/internal/domain/abstractions"
)

type Repository struct {
	PerformedQueries abstractions.PerformedQueriesRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		PerformedQueries: NewPqRepository(db),
	}
}
