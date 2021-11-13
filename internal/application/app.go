package application

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

func (a application) GetPerformedQuery(query string) (*entity.PerformedQuery, error) {
	pq, err := a.repo.PerformedQueriesStore.GetByQuery(query)

	if err != nil {
		return nil, err
	}

	return pq, nil
}

func (a application) CreateNewPerformedQuery(query string) (*int64, error) {
	_, err := a.repo.PerformedQueriesStore.Create(query)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (a application) CreateNewRecipe(recipe *entity.Recipe) error {
	err := a.repo.RecipeStore.Create(recipe)

	if err != nil {
		return err
	}

	return nil
}

func (a application) GetRecipesByQueryId(id int64) ([]entity.Recipe, error) {
	recipes, err := a.repo.RecipeStore.GetByQueryId(id)

	if err != nil {
		return nil, err
	}

	return recipes, nil
}
