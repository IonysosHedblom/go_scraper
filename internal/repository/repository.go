package repository

import (
	"database/sql"

	"github.com/ionysoshedblom/go_scraper/internal/domain/abstractions"
)

type Repository struct {
	PerformedQueriesStore   abstractions.PerformedQueriesStore
	RecipeStore             abstractions.RecipeStore
	IngredientSearchesStore abstractions.IngredientSearchesStore
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		PerformedQueriesStore:   NewPerformedQueriesStore(db),
		RecipeStore:             NewRecipeStore(db),
		IngredientSearchesStore: NewIngredientSearchStore(db),
	}
}
