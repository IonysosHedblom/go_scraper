package handlers

import (
	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"github.com/ionysoshedblom/go_scraper/internal/repository"
)

type inventoryHandler struct {
	repo *repository.Repository
}

func NewInventoryHandler(repo *repository.Repository) *inventoryHandler {
	return &inventoryHandler{repo: repo}
}

func (ih *inventoryHandler) Get(userId string) ([]entity.Item, error) {
	items, err := ih.repo.InventoryStore.GetItems(userId)

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (ih *inventoryHandler) AddItems(items []entity.InventoryItem) error {
	err := ih.repo.InventoryStore.AddItemsToUserInventory(items)

	if err != nil {
		return err
	}

	return nil
}
func (ih *inventoryHandler) Create(userId string) error {
	return nil
}
