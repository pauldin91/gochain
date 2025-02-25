package app

import (
	"net/http"
)

// balanceHandler retrieves the address for a given wallet
// @Summary      Get balance
// @Description  Retrieves the address of a wallet
// @Tags         public-key
// @Produce      json
// @Success      200
// @Router       /public-key [get]
func publicKeyHandler(writer http.ResponseWriter, req *http.Request) {

}
