package api

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type ScraperService interface {
	HandleSource(src *html.Node) ([]entity.Recipe, error)
}
