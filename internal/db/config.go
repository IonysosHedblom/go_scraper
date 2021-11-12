package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func CreateDbDSN() string {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", getEnv("DB_USER"), getEnv("DB_PASSWORD"), getEnv("DB_NAME"))
	return dsn
}
