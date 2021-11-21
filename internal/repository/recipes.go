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

func (r *recipeStore) GetById(id int64) (*entity.Recipe, error) {
	recipe := new(entity.Recipe)

	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(
		ctx,
		"SELECT * FROM recipes WHERE recipe_id = $1", id,
	).Scan(
		&recipe.Id,
		&recipe.Title,
		&recipe.Description,
		&recipe.ImageUrl,
		pq.Array(&recipe.Ingredients),
		&recipe.QueryId,
		&recipe.IngredientSearchId,
		pq.Array(&recipe.Checklist),
		&recipe.Rating,
	)

	if err != nil {
		return nil, err
	}

	return recipe, nil
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
		err := rows.Scan(
			&recipe.Id,
			&recipe.Title,
			&recipe.Description,
			&recipe.ImageUrl,
			pq.Array(&recipe.Ingredients),
			pq.Array(&recipe.Checklist),
			&recipe.Rating,
			&recipe.QueryId,
			&recipe.IngredientSearchId,
		)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, *recipe)
	}

	return recipes, nil
}

func (r *recipeStore) GetByIngredientSearchId(id int64) ([]entity.Recipe, error) {
	var recipes []entity.Recipe

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, "SELECT * FROM recipes WHERE ingredient_search_id = $1", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		recipe := new(entity.Recipe)
		err := rows.Scan(&recipe.Id,
			&recipe.Title,
			&recipe.Description,
			&recipe.ImageUrl,
			pq.Array(&recipe.Ingredients),
			pq.Array(&recipe.Checklist),
			&recipe.Rating,
			&recipe.QueryId,
			&recipe.IngredientSearchId,
		)
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

	_, err = statement.ExecContext(
		ctx,
		recipe.Id,
		recipe.Title,
		recipe.Description,
		recipe.ImageUrl,
		pq.Array(recipe.Ingredients),
		*recipe.QueryId,
	)
	return err
}

func (r *recipeStore) CreateFromIngredients(recipe *entity.Recipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbQuery := "INSERT INTO recipes (recipe_id, title, description, imageurl, ingredients, ingredient_search_id) VALUES ($1, $2, $3, $4, $5, $6)"

	statement, err := r.db.PrepareContext(ctx, dbQuery)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.ExecContext(
		ctx,
		recipe.Id,
		recipe.Title,
		recipe.Description,
		recipe.ImageUrl,
		pq.Array(recipe.Ingredients),
		*recipe.IngredientSearchId,
	)
	return err
}

func (r *recipeStore) Update(ingredients, checklist []string, rating string, recipeId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbQuery := "UPDATE recipes SET ingredients = $1, checklist = $2, rating = $3 WHERE recipe_id = $4"
	statement, err := r.db.PrepareContext(ctx, dbQuery)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.ExecContext(ctx, pq.Array(ingredients), pq.Array(checklist), rating, recipeId)

	return err
}

func (r *recipeStore) UpdateIngredientSearchId(ingredientSearchId *int64, recipeId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbQuery := "UPDATE recipes SET ingredient_search_id = $1 WHERE recipe_id = $2"

	statement, err := r.db.PrepareContext(ctx, dbQuery)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.ExecContext(ctx, *ingredientSearchId, recipeId)

	return err
}

func (r *recipeStore) UpdateQueryId(queryId *int64, recipeId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbQuery := "UPDATE recipes SET query_id = $1 WHERE recipe_id = $2"

	statement, err := r.db.PrepareContext(ctx, dbQuery)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.ExecContext(ctx, *queryId, recipeId)

	return err
}
