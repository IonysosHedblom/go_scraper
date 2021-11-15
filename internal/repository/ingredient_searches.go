package repository

import (
	"database/sql"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
)

type ingredientSearchesStore struct {
	db *sql.DB
}

func NewIngredientSearchStore(db *sql.DB) *ingredientSearchesStore {
	return &ingredientSearchesStore{db: db}
}

func (i *ingredientSearchesStore) GetByIngredients(ingredients []string) (*entity.IngredientSearch, error) {
	return nil, nil
}

func (i *ingredientSearchesStore) GetById(id int) (*entity.IngredientSearch, error) {
	return nil, nil
}

func (i *ingredientSearchesStore) Create(ingredients []string) (*int64, error) {
	return nil, nil
}
