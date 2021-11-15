package application

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/abstractions"
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type application struct {
	scraper abstractions.ScraperPort
}

func NewApplication(scraper abstractions.ScraperPort) *application {
	return &application{
		scraper: scraper,
	}
}

func (a application) CallRecipeResultScraping(src *html.Node) ([]entity.Recipe, error) {
	stringSrc, err := a.scraper.GetRecipeResults(src)

	if err != nil {
		return nil, err
	}

	return stringSrc, nil
}
