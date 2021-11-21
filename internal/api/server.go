package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ionysoshedblom/go_scraper/internal/application/handlers"
	"github.com/ionysoshedblom/go_scraper/internal/domain/abstractions"
)

type api struct {
	app      abstractions.AppPort
	handlers *handlers.Handlers
}

func NewApi(app abstractions.AppPort, handlers *handlers.Handlers) *api {
	return &api{app: app, handlers: handlers}
}

func (a *api) Run() {
	http.HandleFunc("/api/scraper", a.ScraperRouter)
	http.HandleFunc("/api/recipeDetails", a.RecipeDetailsRouter)
	http.HandleFunc("/api/inventory", a.InventoryRouter)
	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
