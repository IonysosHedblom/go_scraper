package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
)

type itemsStore struct {
	db *sql.DB
}

func NewItemsStore(db *sql.DB) *itemsStore {
	return &itemsStore{db: db}
}

func (i *itemsStore) GetItemsByUserId(id string) ([]entity.Item, error) {
	var items []entity.Item

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbQuery := `SELECT (name, imageurl, expire_date, volume_amount, volume_unit, category, quantity) FROM items
	INNER JOIN inventory_items ON items.item_id = inventory_items.item_id
	INNER JOIN inventories ON inventory_items.inventory_id = inventories.inventory_id
	WHERE inventories.user_id = $1`

	rows, err := i.db.QueryContext(ctx, dbQuery, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		item := new(entity.Item)
		err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.ImageUrl,
			&item.ExpireDate,
			&item.VolumeAmount,
			&item.VolumeUnit,
			&item.Category,
			&item.Quantity,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return items, nil
}

func (i *itemsStore) AddItemsToUserInventory(id int64, items []entity.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbQuery := "INSERT INTO inventory_items (item_id, inventory_id, quantity) VALUES ($1, $2, $3)"

	statement, err := i.db.PrepareContext(ctx, dbQuery)

	if err != nil {
		return err
	}

	defer statement.Close()

	for _, item := range items {
		_, err = statement.ExecContext(
			ctx,
			item.Id,
			id,
			item.Quantity,
		)
	}

	return err
}
