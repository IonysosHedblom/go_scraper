package repository

import (
	"context"
	"database/sql"
	"time"
)

type inventoryStore struct {
	db *sql.DB
}

func NewInventoryStore(db *sql.DB) *inventoryStore {
	return &inventoryStore{db: db}
}

func (is *inventoryStore) Create(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbQuery := "INSERT INTO inventories (user_id) VALUES ($1)"

	statement, err := is.db.PrepareContext(ctx, dbQuery)

	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.ExecContext(ctx, id)
	return err
}
