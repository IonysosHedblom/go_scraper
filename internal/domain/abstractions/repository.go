package abstractions

import "github.com/ionysoshedblom/go_scraper/internal/domain/entity"

type Repository interface {
	Get(query string) int
	Create(query string) entity.PerformedQuery
}
