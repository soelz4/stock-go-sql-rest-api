package stock

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"stock-go-sql-rest-api/src/types"
	"stock-go-sql-rest-api/src/utils"
)

func (h *Handler) handleCreateStock(w http.ResponseWriter, r *http.Request) {
	var stock types.Stock

	err := utils.ParseJSON(r, &stock)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	id, err := h.store.InsertStock(stock)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	stock.StockID = int64(id)

	utils.WriteJSON(w, http.StatusOK, stock)
}

func (h *Handler) handleGetStockByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing stock id"))
	}

	stockID, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid stock id"))
	}

	stock, err := h.store.GetStockByID(stockID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, stock)
}

func (h *Handler) handleGetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := h.store.GetAllStocks()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, stocks)
}

func (h *Handler) handleUpdateStockByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing stock id"))
	}

	var stock types.Stock
	err := utils.ParseJSON(r, &stock)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	stockID, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid stock id"))
	}

	stock.StockID = int64(stockID)

	rowsAffected, err := h.store.UpdateStockByID(stock)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		map[string]interface{}{"Stock": stock, "Rows Affected": rowsAffected},
	)
}

func (h *Handler) handleDeleteStockByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing stock id"))
	}

	stockID, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid stock id"))
	}

	rowsAffected, err := h.store.DeleteStockByID(stockID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, map[string]int64{"Rows Affected": rowsAffected})
}
