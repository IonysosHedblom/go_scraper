package handlers

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"github.com/ionysoshedblom/go_scraper/internal/repository"
)

type recipeHandler struct {
	repo *repository.Repository
}

func NewRecipeHandler(repo *repository.Repository) *recipeHandler {
	return &recipeHandler{repo: repo}
}

func (rh *recipeHandler) CreateNewRecipe(recipe *entity.Recipe) error {
	err := rh.repo.RecipeStore.Create(recipe)

	if err != nil {
		return err
	}

	return nil
}

func (rh *recipeHandler) CreateNewRecipeFromIngredients(recipe *entity.Recipe) error {
	err := rh.repo.RecipeStore.CreateFromIngredients(recipe)

	if err != nil {
		return err
	}

	return nil
}

func (rh *recipeHandler) GetRecipeById(id int64) (*entity.Recipe, error) {
	recipe, err := rh.repo.RecipeStore.GetById(id)

	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (rh *recipeHandler) GetRecipesByQueryId(id int64) ([]entity.Recipe, error) {
	recipes, err := rh.repo.RecipeStore.GetByQueryId(id)

	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (rh *recipeHandler) GetRecipesByIngredientSearchId(id int64) ([]entity.Recipe, error) {
	recipes, err := rh.repo.RecipeStore.GetByIngredientSearchId(id)

	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (rh *recipeHandler) UpdateRecipeIngredientSearchId(ingredientSearchId *int64, recipeId int64) error {
	if err := rh.repo.RecipeStore.UpdateIngredientSearchId(ingredientSearchId, recipeId); err != nil {
		return err
	}

	return nil
}

func (rh *recipeHandler) UpdateRecipeQueryId(queryId *int64, recipeId int64) error {
	if err := rh.repo.RecipeStore.UpdateQueryId(queryId, recipeId); err != nil {
		return err
	}

	return nil
}
