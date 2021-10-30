package api

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type Application struct {
	scraper ScraperService
}

func NewApplication(scraper ScraperService) *Application {
	return &Application{scraper: scraper}
}

func (a Application) HandleSource(src *html.Node) ([]entity.Recipe, error) {
	stringSrc, err := a.scraper.HandleSource(src)

	if err != nil {
		return nil, err
	}

	return stringSrc, nil
}
