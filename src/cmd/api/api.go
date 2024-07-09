package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"stock-go-sql-rest-api/src/services/stock"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{address: address, db: db}
}

func (s *APIServer) Run() error {
	// Gorilla Mux Router
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// Stock Handler
	stockStore := stock.NewStore(s.db)
	stockHandler := stock.NewHandler(stockStore)
	stockHandler.RegisterRoutes(subRouter)
	log.Println("Listening On", s.address)
	return http.ListenAndServe(s.address, router)
}
