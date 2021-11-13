package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
)

type recipes struct {
	db *sql.DB
}

func NewRecipeStore(db *sql.DB) *recipes {
	return &recipes{
		db: db,
	}
}

func (r *recipes) GetByQueryId(id int) (*entity.Recipe, error) {
	recipe := new(entity.Recipe)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, "SELECT * FROM recipes WHERE query_id = $1", id).Scan(&recipe.Title, &recipe.Description, &recipe.ImageUrl, &recipe.Ingredients)

	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (r *recipes) Create(recipe *entity.Recipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbQuery := "INSERT INTO recipes (title, description, imageurl, ingredients) VALUES ($1, $2, $3, $4)"
	statement, err := r.db.PrepareContext(ctx, dbQuery)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.ExecContext(ctx, recipe.Title, recipe.Description, recipe.ImageUrl, recipe.Ingredients)
	return err
}
