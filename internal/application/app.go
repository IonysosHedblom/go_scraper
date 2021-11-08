package app

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/abstractions"
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type Application struct {
	scraper abstractions.ScraperPort
}

func NewApplication(scraper abstractions.ScraperPort) *Application {
	return &Application{scraper: scraper}
}

func (a Application) GetByQueryHandler(src *html.Node) []entity.Recipe {
	stringSrc := a.scraper.GetRecipeResults(src)

	return stringSrc
}
