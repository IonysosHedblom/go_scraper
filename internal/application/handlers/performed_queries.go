package handlers

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"github.com/ionysoshedblom/go_scraper/internal/repository"
)

type pqHandler struct {
	repo *repository.Repository
}

func NewPqHandler(repo *repository.Repository) *pqHandler {
	return &pqHandler{repo: repo}
}

func (pqh *pqHandler) GetPerformedQuery(query string) (*entity.PerformedQuery, error) {
	pq, err := pqh.repo.PerformedQueriesStore.GetByQuery(query)

	if err != nil {
		return nil, err
	}

	return pq, nil
}

func (pqh *pqHandler) CreateNewPerformedQuery(query string) (*int64, error) {
	queryId, err := pqh.repo.PerformedQueriesStore.Create(query)

	if err != nil {
		return nil, err
	}

	return queryId, nil
}
