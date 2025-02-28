package app

import (
	"encoding/json"
	"net/http"
)

// balanceHandler retrieves the address for a given wallet
// @Summary      Get balance
// @Description  Retrieves the address of a wallet
// @Tags         public-key
// @Produce      json
// @Success      200 {object} WalletDto
// @Router       /public-key [get]
func (s *HttpServer) publicKeyHandler(writer http.ResponseWriter, req *http.Request) {
	wallet := WalletDto{
		Address: s.wallet.Address,
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(wallet)
}
