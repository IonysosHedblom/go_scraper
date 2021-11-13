package abstractions

import "github.com/ionysoshedblom/go_scraper/internal/domain/entity"

type PerformedQueriesStore interface {
	GetByQuery(query string) (*entity.PerformedQuery, error)
	GetById(id int) (*entity.PerformedQuery, error)
	Create(query string) error
}

type RecipeStore interface {
	GetByQueryId(id int) (*entity.Recipe, error)
	Create(recipe *entity.Recipe) error
}
