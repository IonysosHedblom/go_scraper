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

var qid int64 = 5

var recipe = &entity.Recipe{
	Id:          5,
	Title:       "test",
	Description: "testDesc",
	ImageUrl:    "https://testimgurl.jpg",
	Ingredients: []string{"test1", "test2"},
	QueryId:     &qid,
}

func NewRecipeRepositoryMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatalf("an error %s was not expected when opening a stub db connection", err)
	}

	return db, mock
}

func TestGetByQueryId(t *testing.T) {
	db, mock := NewPqRepositoryMock()

	repo := &recipeStore{db}

	defer func() {
		repo.db.Close()
	}()

	dbQuery := "SELECT \\* FROM recipes WHERE query_id = \\$1"

	rows := sqlmock.NewRows(
		[]string{
			"recipe_id",
			"title",
			"description",
			"imageurl",
			"ingredients",
			"query_id",
		}).AddRow(recipe.Id, recipe.Title, recipe.Description, recipe.ImageUrl, pq.Array(recipe.Ingredients), recipe.QueryId)

	mock.ExpectQuery(dbQuery).WithArgs(recipe.QueryId).WillReturnRows(rows)

	res, err := repo.GetByQueryId(*recipe.QueryId)
	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestGetByQueryIdError(t *testing.T) {
	db, mock := NewPqRepositoryMock()

	repo := &recipeStore{db}

	defer func() {
		repo.db.Close()
	}()

	dbQuery := "SELECT * FROM recipes WHERE query_id = \\$1"

	rows := sqlmock.NewRows(
		[]string{
			"recipe_id",
			"title",
			"description",
			"imageurl",
			"ingredients",
			"query_id",
		})

	mock.ExpectQuery(dbQuery).WithArgs(recipe.QueryId).WillReturnRows(rows)

	res, err := repo.GetByQueryId(*recipe.QueryId)
	assert.Empty(t, res)
	assert.Error(t, err)
}

func TestCreate(t *testing.T) {
	db, mock := NewPqRepositoryMock()

	repo := &recipeStore{db}

	defer func() {
		repo.db.Close()
	}()

	dbQuery := "INSERT INTO recipes \\(recipe_id, title, description, imageurl, ingredients, query_id\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6\\)"

	prep := mock.ExpectPrepare(dbQuery)
	prep.ExpectExec().WithArgs(recipe.Id, recipe.Title, recipe.Description, recipe.ImageUrl, pq.Array(recipe.Ingredients), recipe.QueryId).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Create(recipe)
	assert.NoError(t, err)
}
