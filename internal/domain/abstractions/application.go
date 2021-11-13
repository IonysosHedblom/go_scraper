package abstractions

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type AppPort interface {
	CallRecipeResultScraping(*html.Node) []entity.Recipe
	GetPerformedQuery(query string) (*entity.PerformedQuery, error)
	CreateNewPerformedQuery(query string) (*int64, error)
	CreateNewRecipe(recipe *entity.Recipe) error
	GetRecipesByQueryId(id int64) ([]entity.Recipe, error)
}
