package repository

import "database/sql"

type ingredientSearchesStore struct {
	db *sql.DB
}

func NewIngredientSearchStore(db *sql.DB) *ingredientSearchesStore {
	return &ingredientSearchesStore{db: db}
}
