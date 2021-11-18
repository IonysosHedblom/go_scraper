package abstractions

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type ScraperPort interface {
	GetRecipeResults(src *html.Node) ([]entity.Recipe, error)
	GetRecipeDetails(n *html.Node) string
}
