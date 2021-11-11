package repository

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

var performed_query = &entity.PerformedQuery{
	Id:    1,
	Query: "pasta",
}

func NewRepositoryMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error %s was not expected when opening a stub db connection", err)
	}

	return db, mock
}

func TestFindByQuery(t *testing.T) {
	db, mock := NewRepositoryMock()

	repo := &repository{db}

	defer func() {
		repo.db.Close()
	}()

	dbQuery := "SELECT query_id, query FROM performed_queries WHERE query = \\$1"

	rows := sqlmock.NewRows([]string{"query_id", "query"}).AddRow(performed_query.Id, performed_query.Query)

	mock.ExpectQuery(dbQuery).WithArgs(performed_query.Query).WillReturnRows(rows)

	pq, err := repo.GetByQuery(performed_query.Query)
	assert.NotNil(t, pq)
	assert.NoError(t, err)
}

func TestFindByQueryError(t *testing.T) {
	db, mock := NewRepositoryMock()

	repo := &repository{db}

	defer func() {
		repo.db.Close()
	}()

	dbQuery := "SELECT query_id, query FROM performed_queries WHERE query = \\$1"

	rows := sqlmock.NewRows([]string{"query_id", "query"})

	mock.ExpectQuery(dbQuery).WithArgs(performed_query.Query).WillReturnRows(rows)

	pq, err := repo.GetByQuery(performed_query.Query)
	assert.Empty(t, pq)
	assert.Error(t, err)
}

func TestFindById(t *testing.T) {
	db, mock := NewRepositoryMock()

	repo := &repository{db}

	defer func() {
		repo.db.Close()
	}()

	dbQuery := "SELECT query_id, query FROM performed_queries WHERE query_id = \\$1"

	rows := sqlmock.NewRows([]string{"query_id", "query"}).AddRow(performed_query.Id, performed_query.Query)

	mock.ExpectQuery(dbQuery).WithArgs(performed_query.Id).WillReturnRows(rows)

	pq, err := repo.GetById(performed_query.Id)
	assert.NotNil(t, pq)
	assert.NoError(t, err)
}

func TestFindByIdError(t *testing.T) {
	db, mock := NewRepositoryMock()

	repo := &repository{db}

	defer func() {
		repo.db.Close()
	}()

	dbQuery := "SELECT query_id, query FROM performed_queries WHERE query_id = \\$1"

	rows := sqlmock.NewRows([]string{"query_id", "query"})

	mock.ExpectQuery(dbQuery).WithArgs(performed_query.Id).WillReturnRows(rows)

	pq, err := repo.GetById(performed_query.Id)
	assert.Empty(t, pq)
	assert.Error(t, err)
}

func TestCreate(t *testing.T) {
	db, mock := NewRepositoryMock()

	repo := &repository{db}

	defer func() {
		repo.db.Close()
	}()

	dbQuery := "INSERT INTO performed_queries \\(query\\) VALUES \\(\\$1\\)"

	preparation := mock.ExpectPrepare(dbQuery)
	preparation.ExpectExec().WithArgs(performed_query.Query).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Create(performed_query.Query)

	assert.NoError(t, err)
}

func TestCreateError(t *testing.T) {
	db, mock := NewRepositoryMock()

	repo := &repository{db}

	defer func() {
		repo.db.Close()
	}()

	badDbQuery := "INSERT INTO performed_query \\(query\\) VALUES \\(\\$1\\)"

	preparation := mock.ExpectPrepare(badDbQuery)
	preparation.ExpectExec().WithArgs(performed_query.Query).WillReturnResult(sqlmock.NewResult(0, 0))
	err := repo.Create(performed_query.Query)
	assert.Error(t, err)
}
