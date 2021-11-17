package abstractions

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type AppPort interface {
	CallRecipeResultScraping(*html.Node) ([]entity.Recipe, error)
	CallRecipeDetailsScraping(src *html.Node) entity.RecipeDetails
}
