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

func (a application) CallRecipeResultScraping(src *html.Node) ([]entity.Recipe, error) {
	stringSrc, err := a.scraper.GetRecipeResults(src)

	if err != nil {
		return nil, err
	}

	return stringSrc, nil
}

func (a application) GetPerformedQuery(query string) (*entity.PerformedQuery, error) {
	pq, err := a.repo.PerformedQueriesStore.GetByQuery(query)

	if err != nil {
		return nil, err
	}

	return pq, nil
}

func (a application) CreateNewPerformedQuery(query string) (*int64, error) {
	queryId, err := a.repo.PerformedQueriesStore.Create(query)

	if err != nil {
		return nil, err
	}

	return queryId, nil
}

func (a application) CreateNewRecipe(recipe *entity.Recipe) error {
	err := a.repo.RecipeStore.Create(recipe)

	if err != nil {
		return err
	}

	return nil
}

func (a application) CreateNewRecipeFromIngredients(recipe *entity.Recipe) error {
	err := a.repo.RecipeStore.CreateFromIngredients(recipe)

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

func (a application) GetIngredientSearch(ingredients []string) (*entity.IngredientSearch, error) {
	ingS, err := a.repo.IngredientSearchesStore.GetByIngredients(ingredients)

	if err != nil {
		return nil, err
	}

	return ingS, err
}

func (a application) CreateIngredientSearch(ingredients []string) (*int64, error) {
	id, err := a.repo.IngredientSearchesStore.Create(ingredients)

	if err != nil {
		return nil, err
	}

	return id, err
}

func (a application) GetRecipesByIngredientSearchId(id int64) ([]entity.Recipe, error) {
	recipes, err := a.repo.RecipeStore.GetByIngredientSearchId(id)

	if err != nil {
		return nil, err
	}

	return recipes, nil
}
