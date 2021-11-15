package abstractions

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type AppPort interface {
	CallRecipeResultScraping(*html.Node) ([]entity.Recipe, error)
	GetPerformedQuery(query string) (*entity.PerformedQuery, error)
	CreateNewPerformedQuery(query string) (*int64, error)
	CreateNewRecipe(recipe *entity.Recipe) error
	CreateNewRecipeFromIngredients(recipe *entity.Recipe) error
	GetRecipeById(id int64) (*entity.Recipe, error)
	GetRecipesByQueryId(id int64) ([]entity.Recipe, error)
	GetRecipesByIngredientSearchId(id int64) ([]entity.Recipe, error)
	GetIngredientSearch(ingredients []string) (*entity.IngredientSearch, error)
	CreateIngredientSearch(ingredients []string) (*int64, error)
}
