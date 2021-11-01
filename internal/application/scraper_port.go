package app

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type ScraperPort interface {
	HandleSource(src *html.Node) ([]entity.Recipe, error)
}
