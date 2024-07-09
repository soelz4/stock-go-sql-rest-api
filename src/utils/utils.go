package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Parse JSON
func ParseJSON(r *http.Request, x interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	} else {
		return json.NewDecoder(r.Body).Decode(x)
	}
}

// Write JSON
func WriteJSON(w http.ResponseWriter, status int, x interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(x)
}

// Write Error
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
