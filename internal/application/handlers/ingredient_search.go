package handlers

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"github.com/ionysoshedblom/go_scraper/internal/repository"
)

type ingredientSearchHandler struct {
	repo *repository.Repository
}

func NewIngredientSearchHandler(repo *repository.Repository) *ingredientSearchHandler {
	return &ingredientSearchHandler{repo: repo}
}

func (ish *ingredientSearchHandler) GetIngredientSearch(ingredients []string) (*entity.IngredientSearch, error) {
	ingS, err := ish.repo.IngredientSearchesStore.GetByIngredients(ingredients)

	if err != nil {
		return nil, err
	}

	return ingS, err
}

func (ish *ingredientSearchHandler) CreateIngredientSearch(ingredients []string) (*int64, error) {
	id, err := ish.repo.IngredientSearchesStore.Create(ingredients)

	if err != nil {
		return nil, err
	}

	return id, err
}
