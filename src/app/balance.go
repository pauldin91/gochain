package app

import (
	"encoding/json"
	"net/http"
)

type BalanceResponse struct {
	Message string `json:"message"`
}

func balanceHandler(w http.ResponseWriter, req *http.Request) {

	response := BalanceResponse{
		Message: "balance is 0",
	}
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode response as JSON and write to response writer
	json.NewEncoder(w).Encode(response)
}
