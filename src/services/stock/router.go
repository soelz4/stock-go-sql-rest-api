package stock

import (
	"github.com/gorilla/mux"

	"stock-go-sql-rest-api/src/types"
)

type Handler struct {
	store types.StockStore
}

func NewHandler(store types.StockStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/stock/{id}", h.handleGetStockByID).Methods("GET", "OPTIONS")
	router.HandleFunc("/stocks", h.handleGetAllStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/newstock", h.handleCreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/stock/{id}", h.handleUpdateStockByID).Methods("PUT", "OPTIONS")
	router.HandleFunc("/deletestock/{id}", h.handleDeleteStockByID).Methods("DELETE", "OPTIONS")
}
