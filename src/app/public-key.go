package app

import (
	"encoding/json"
	"net/http"
)

// retrieves the address for a given wallet
// @Summary      Get public key
// @Description  Retrieves the address of a wallet
// @Tags         public-key
// @Produce      json
// @Success      200 {object} WalletDto
// @Router       /public-key [get]
func (s *Peer) publicKeyHandler(writer http.ResponseWriter, req *http.Request) {
	wallet := WalletDto{
		Address: s.wallet.Address,
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(wallet)
}
