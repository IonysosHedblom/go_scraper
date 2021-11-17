package abstractions

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
)

type RecipeHandler interface {
	CreateNewRecipe(recipe *entity.Recipe) error
	CreateNewRecipeFromIngredients(recipe *entity.Recipe) error
	GetRecipeById(id int64) (*entity.Recipe, error)
	GetRecipesByQueryId(id int64) ([]entity.Recipe, error)
	GetRecipesByIngredientSearchId(id int64) ([]entity.Recipe, error)
	UpdateRecipeQueryId(queryId *int64, recipeId int64) error
	UpdateRecipeIngredientSearchId(ingredientSearchId *int64, recipeId int64) error
	UpdateIngredientsAndChecklist(ingredients, checklist []string, recipeId int64) error
}

type PerformedQueryHandler interface {
	GetPerformedQuery(query string) (*entity.PerformedQuery, error)
	CreateNewPerformedQuery(query string) (*int64, error)
}

type IngredientSearchHandler interface {
	GetIngredientSearch(ingredients []string) (*entity.IngredientSearch, error)
	CreateIngredientSearch(ingredients []string) (*int64, error)
}
