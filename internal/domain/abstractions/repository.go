package abstractions

import "github.com/ionysoshedblom/go_scraper/internal/domain/entity"

type Repository interface {
	GetPerformedQueryByQuery(query string) (*entity.PerformedQuery, error)
	GetPerformedQueryById(id int) (*entity.PerformedQuery, error)
	CreatePerformedQuery(query string) error
}
