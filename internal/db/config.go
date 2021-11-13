package db

import (
	"fmt"
)

func CreateDbDSN() string {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	return dsn
}
