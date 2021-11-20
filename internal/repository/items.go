package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type itemsStore struct {
	db *sql.DB
}

func NewItemsStore(db *sql.DB) *itemsStore {
	return &itemsStore{db: db}
}

func (i *itemsStore) InsertRows() {
	// jMap := getItemMap()
	// ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// defer cancel()

	// for _, value := range jMap {
	// 	// To insert to each row
	// 	// fmt.Println(value["Name"])
	// 	dbQuery := "INSERT INTO items (name) VALUES ($1)"
	// 	statement, err := r.db.PrepareContext(ctx, dbQuery)
	// }
}

func getItemMap() []map[string]string {
	jsonFile, err := os.Open("localization.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var jMap []map[string]string
	err = json.Unmarshal(byteValue, &jMap)

	if err != nil {
		fmt.Println(err)
	}

	return jMap
}
