package stock

import (
	"database/sql"
	"log"

	"stock-go-sql-rest-api/src/types"
)

// Insert Stock into DataBase
func (s *Store) InsertStock(stock types.Stock) (int64, error) {
	// Inserted ID
	var id int64

	// Query
	query := "INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING stockid"
	err := s.db.QueryRow(query, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		return id, err
	} else {
		return id, err
	}
}

// GET Stock By ID
func (s *Store) GetStockByID(id int) (*types.Stock, error) {
	// Query
	query := "SELECT * FROM stocks WHERE stockid = $1"
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	stock := new(types.Stock)
	for rows.Next() {
		stock, err = scanRowIntoStock(rows)
		if err != nil {
			return nil, err
		}
	}

	return stock, err
}

// GET All Stocks
func (s *Store) GetAllStocks() ([]*types.Stock, error) {
	// Query
	query := "SELECT * FROM stocks"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	stocks := make([]*types.Stock, 0)

	for rows.Next() {
		stock, err := scanRowIntoStock(rows)
		if err != nil {
			return nil, err
		}

		stocks = append(stocks, stock)
	}

	return stocks, err
}

// Update Stock By ID
func (s *Store) UpdateStockByID(stock types.Stock) (int64, error) {
	// Query
	query := "UPDATE stocks SET name= $2, price= $3, company= $4 WHERE stockid = $1"
	result, err := s.db.Exec(query, stock.StockID, stock.Name, stock.Price, stock.Company)
	if err != nil {
		log.Fatalf("Unable to Execute the Query, %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Error While Checking the Affected Rows, %v", err)
	}

	return rowsAffected, err
}

// Delete Stock By ID
func (s *Store) DeleteStockByID(id int) (int64, error) {
	// Query
	query := "DELETE FROM stocks WHERE stockid = $1"
	result, err := s.db.Exec(query, id)
	if err != nil {
		log.Fatalf("Unable to Execute the Query, %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Error While Checking the Affected Rows, %v", err)
	}

	return rowsAffected, err
}

// PUT DataBase Rows into Stock Struct Type
func scanRowIntoStock(rows *sql.Rows) (*types.Stock, error) {
	stock := new(types.Stock)

	err := rows.Scan(
		&stock.StockID,
		&stock.Name,
		&stock.Price,
		&stock.Company,
	)

	if err != nil {
		return nil, err
	} else {
		return stock, err
	}
}
