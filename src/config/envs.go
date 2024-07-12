package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DataBase URL
var DB_URL = initConfig()

// GET DataBase URL from .env File
func initConfig() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env File")
	}

	dbURL := getEnv("POSTGRES_URL")
	return dbURL
}

// GET Env Variables Func
func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	} else {
		return "DB URL is not Valid"
	}
}
