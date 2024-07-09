package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

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

	// Migrate PostgreSQL Driver
	driver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://src/cmd/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]

	// UP
	if cmd == "up" {
		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	// DOWN
	if cmd == "down" {
		err = m.Down()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
