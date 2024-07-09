package stock

import (
	"database/sql"
	"log"

	"stock-go-sql-rest-api/src/types"
)

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

func (s *Store) GetAllStocks() ([]*types.Stock, error) {
	// SQL Query
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
