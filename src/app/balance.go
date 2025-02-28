package app

import (
	"encoding/json"
	"net/http"
)

// balanceHandler retrieves the balance for a given wallet
// @Summary      Get balance
// @Description  Retrieves the balance of a wallet
// @Tags         balance
// @Produce      json
// @Success      200  {array}   BalanceResponse
// @Router       /balance [get]
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
