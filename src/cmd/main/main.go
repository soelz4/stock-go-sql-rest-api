package main

import (
	"fmt"
	"log"

	"stock-go-sql-rest-api/src/cmd/api"
	"stock-go-sql-rest-api/src/config"
	"stock-go-sql-rest-api/src/db"
)

func main() {
	// DataBase URL (From .env File)
	dbURL := config.DB_URL

	// Initialize and Start Connection with DataBase
	database, err := db.NewPostgreSQLStorage(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	db.InitPostgreSQLStorage(database)

	server := api.NewAPIServer(fmt.Sprintf(":%s", "9010"), database)
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
