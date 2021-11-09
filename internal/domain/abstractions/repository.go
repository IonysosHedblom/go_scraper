package abstractions

import "github.com/ionysoshedblom/go_scraper/internal/domain/entity"

type PerformedQueriesRepository interface {
	Get(query string) (id int)
	Create(query string) entity.PerformedQuery
}
