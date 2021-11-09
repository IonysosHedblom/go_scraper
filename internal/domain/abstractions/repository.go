package abstractions

import "github.com/ionysoshedblom/go_scraper/internal/domain/entity"

type Repository interface {
	GetByQuery(query string) (*entity.PerformedQuery, error)
	GetById(id int) (*entity.PerformedQuery, error)
	Create(query string) error
}
