package types

// User Schema of the Stocks Table
type Stock struct {
	StockID int64  `json:"stockid"`
	Name    string `json:"name"`
	Price   int64  `json:"price"`
	Company string `json:"company"`
}

// Stock Storage
type StockStore interface {
	InsertStock(Stock) (int64, error)
	GetStockByID(id int) (*Stock, error)
	GetAllStocks() ([]*Stock, error)
	UpdateStockByID(Stock) (int64, error)
	DeleteStockByID(id int) (int64, error)
}
