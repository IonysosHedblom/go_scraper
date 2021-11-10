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
