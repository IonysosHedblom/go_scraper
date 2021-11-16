package repository

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var ingredient_search = &entity.IngredientSearch{
	Id:          1,
	Ingredients: []string{"test_ingredient", "test_ingredient2"},
}

func NewIngredientSearchRepoMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatalf("an error %s was not expected when opening a stub db connection", err)
	}

	return db, mock
}

func TestGetByIngredients(t *testing.T) {
	db, mock := NewPqRepositoryMock()

	repo := &ingredientSearchesStore{db}

	defer func() {
		repo.db.Close()
	}()

	dbQuery := "SELECT ingredient_search_id, ingredients FROM ingredient_searches WHERE ingredients \\@\\> \\$1 AND ingredients \\<\\@ \\$1"

	rows := sqlmock.NewRows([]string{"ingredient_search_id", "ingredients"}).AddRow(ingredient_search.Id, pq.Array(ingredient_search.Ingredients))

	mock.ExpectQuery(dbQuery).WithArgs(pq.Array(ingredient_search.Ingredients)).WillReturnRows(rows)

	ingredientSearch, err := repo.GetByIngredients(ingredient_search.Ingredients)
	assert.NotNil(t, ingredientSearch)
	assert.NoError(t, err)
}

func TestGetByIngredientsError(t *testing.T) {
	db, mock := NewPqRepositoryMock()

	repo := &ingredientSearchesStore{db}

	defer func() {
		repo.db.Close()
	}()

	dbQuery := "SELECT ingredient_search_id, ingredients FROM ingredient_searches WHERE ingredients \\@\\> \\$1 AND ingredients \\<\\@ \\$1"

	rows := sqlmock.NewRows([]string{"ingredient_search_id", "ingredients"})

	mock.ExpectQuery(dbQuery).WithArgs(pq.Array(ingredient_search.Ingredients)).WillReturnRows(rows)

	ingredientSearch, err := repo.GetByIngredients(ingredient_search.Ingredients)
	assert.Empty(t, ingredientSearch)
	assert.Error(t, err)
}
