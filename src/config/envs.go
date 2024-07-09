package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB_URL = initConfig()

func initConfig() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env File")
	}

	dbURL := getEnv("POSTGRES_URL")
	return dbURL
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	} else {
		return "DB URL is not Valid"
	}
}
