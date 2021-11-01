package app

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type Application struct {
	scraper ScraperPort
}

func NewApplication(scraper ScraperPort) *Application {
	return &Application{scraper: scraper}
}

func (a Application) HandleSource(src *html.Node) ([]entity.Recipe, error) {
	stringSrc, err := a.scraper.HandleSource(src)

	if err != nil {
		return nil, err
	}

	return stringSrc, nil
}
