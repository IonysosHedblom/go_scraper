package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"github.com/lib/pq"
)

type recipeStore struct {
	db *sql.DB
}

func NewRecipeStore(db *sql.DB) *recipeStore {
	return &recipeStore{
		db: db,
	}
}

func (r *recipeStore) GetByQueryId(id int64) ([]entity.Recipe, error) {
	var recipes []entity.Recipe

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, "SELECT * FROM recipes WHERE query_id = $1", id)

	for rows.Next() {
		recipe := new(entity.Recipe)
		err := rows.Scan(&recipe.Title, &recipe.Description, &recipe.ImageUrl, &recipe.Ingredients)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, *recipe)
	}

	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (r *recipeStore) Create(recipe *entity.Recipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbQuery := "INSERT INTO recipes (title, description, imageurl, ingredients) VALUES ($1, $2, $3, $4)"
	statement, err := r.db.PrepareContext(ctx, dbQuery)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.ExecContext(ctx, recipe.Title, recipe.Description, recipe.ImageUrl, pq.Array(recipe.Ingredients))
	return err
}
