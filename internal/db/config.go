package db

import (
	"fmt"
)

// func getEnv(key string) string {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		log.Fatalf("Error loading .env file")
// 	}

// 	return os.Getenv(key)
// }

func CreateDbDSN() string {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	return dsn
}
