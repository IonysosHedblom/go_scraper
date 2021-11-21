package handlers

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/abstractions"
	"github.com/ionysoshedblom/go_scraper/internal/repository"
)

type Handlers struct {
	RecipeHandler           abstractions.RecipeHandler
	PqHandler               abstractions.PerformedQueryHandler
	IngredientSearchHandler abstractions.IngredientSearchHandler
	InventoryHandler        abstractions.InventoryHandler
}

func NewHandlers(repo *repository.Repository) *Handlers {
	return &Handlers{
		RecipeHandler:           NewRecipeHandler(repo),
		PqHandler:               NewPqHandler(repo),
		IngredientSearchHandler: NewIngredientSearchHandler(repo),
		InventoryHandler:        NewInventoryHandler(repo),
	}
}
