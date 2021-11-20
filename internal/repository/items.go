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
