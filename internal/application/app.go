package app

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/abstractions"
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"github.com/ionysoshedblom/go_scraper/internal/repository"
	"golang.org/x/net/html"
)

type application struct {
	scraper abstractions.ScraperPort
	repo    *repository.Repository
}

func NewApplication(scraper abstractions.ScraperPort, repo *repository.Repository) *application {
	return &application{scraper: scraper, repo: repo}
}

func (a application) CallRecipeResultScraping(src *html.Node) []entity.Recipe {
	stringSrc := a.scraper.GetRecipeResults(src)

	return stringSrc
}
