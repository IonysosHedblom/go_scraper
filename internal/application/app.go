package app

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/abstractions"
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type application struct {
	scraper abstractions.ScraperPort
}

func NewApplication(scraper abstractions.ScraperPort) *application {
	return &application{scraper: scraper}
}

func (a application) CallRecipeResultScraping(src *html.Node) []entity.Recipe {
	stringSrc := a.scraper.GetRecipeResults(src)

	return stringSrc
}
