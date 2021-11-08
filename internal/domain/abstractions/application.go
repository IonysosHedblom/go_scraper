package abstractions

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type AppPort interface {
	GetByQueryHandler(*html.Node) []entity.Recipe
}
