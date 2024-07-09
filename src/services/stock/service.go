package stock

import (
	"database/sql"
	"log"

	"stock-go-sql-rest-api/src/types"
)

func (s *Store) InsertStock(stock types.Stock) error {
	_, err := s.db.Exec(
		"INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3)",
		stock.Name,
		stock.Price,
		stock.Company,
	)

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s *Store) GetStockByID(id int) (*types.Stock, error) {
	rows, err := s.db.Query("SELECT * FROM stocks WHERE stockid = $1", id)
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

func (s *Store) UpdateStockByID(id int, stock types.Stock) (int64, error) {
	result, err := s.db.Exec(
		"UPDATE stocks SET name= $1, price= $2, company= $3 WHERE stockid = $4",
		stock.Name,
		stock.Price,
		stock.Company,
		stock.StockID,
	)
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
	result, err := s.db.Exec("DELETE FROM stocks WHERE stockid = $1", id)
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
