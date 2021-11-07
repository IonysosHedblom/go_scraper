package domain

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type ApiPort interface {
	GetByQueryHandler(*html.Node) []entity.Recipe
}
