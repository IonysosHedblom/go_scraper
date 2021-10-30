package domain

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type ApiService interface {
	HandleSource(*html.Node) ([]entity.Recipe, error)
}
