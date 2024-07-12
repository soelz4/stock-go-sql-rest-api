package stock

import "database/sql"

// Stock Storage
type Store struct {
	db *sql.DB
}

// Return New Stock Storage
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}
