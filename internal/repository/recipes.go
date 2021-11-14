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

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		recipe := new(entity.Recipe)
		err := rows.Scan(&recipe.Id, &recipe.Title, &recipe.Description, &recipe.ImageUrl, pq.Array(&recipe.Ingredients), &recipe.QueryId)
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

	dbQuery := "INSERT INTO recipes (recipe_id, title, description, imageurl, ingredients, query_id) VALUES ($1, $2, $3, $4, $5, $6)"
	statement, err := r.db.PrepareContext(ctx, dbQuery)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.ExecContext(ctx, recipe.Id, recipe.Title, recipe.Description, recipe.ImageUrl, pq.Array(recipe.Ingredients), recipe.QueryId)
	return err
}
