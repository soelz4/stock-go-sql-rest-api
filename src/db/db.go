package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgreSQLStorage(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func InitPostgreSQLStorage(db *sql.DB) {
	// Start DataBase Connection
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DataBase Succecfully Connected")
}
