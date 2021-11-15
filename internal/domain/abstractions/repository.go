package abstractions

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
)

type PerformedQueriesStore interface {
	GetByQuery(query string) (*entity.PerformedQuery, error)
	GetById(id int) (*entity.PerformedQuery, error)
	Create(query string) (*int64, error)
}

type RecipeStore interface {
	GetByQueryId(id int64) ([]entity.Recipe, error)
	Create(recipe *entity.Recipe) error
}

type IngredientSearchesStore interface {
	GetByIngredients(ingredients []string) (*entity.IngredientSearch, error)
	GetById(id int) (*entity.IngredientSearch, error)
	Create(ingredients []string) (*int64, error)
}
