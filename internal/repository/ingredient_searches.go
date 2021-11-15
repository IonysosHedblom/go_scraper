package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"github.com/lib/pq"
)

type ingredientSearchesStore struct {
	db *sql.DB
}

func NewIngredientSearchStore(db *sql.DB) *ingredientSearchesStore {
	return &ingredientSearchesStore{db: db}
}

func (i *ingredientSearchesStore) GetByIngredients(ingredients []string) (*entity.IngredientSearch, error) {
	ingredientSearch := new(entity.IngredientSearch)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := i.db.QueryRowContext(ctx, "SELECT ingredient_search_id, ingredients FROM ingredient_searches WHERE ingredients @> $1 AND ingredients <@ $1", pq.Array(ingredients)).Scan(&ingredientSearch.Id, pq.Array(&ingredientSearch.Ingredients))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return ingredientSearch, nil
}

func (i *ingredientSearchesStore) GetById(id int) (*entity.IngredientSearch, error) {
	return nil, nil
}

func (i *ingredientSearchesStore) Create(ingredients []string) (*int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id int64
	dbquery := "INSERT INTO ingredient_searches (ingredients) VALUES ($1) RETURNING ingredient_search_id"
	stmt, err := i.db.PrepareContext(ctx, dbquery)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = stmt.QueryRowContext(ctx, pq.Array(ingredients)).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &id, nil
}
