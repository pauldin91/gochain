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

// peerDiscoveryHandler retrieves hello from world
// @Summary      Get hello world
// @Description  Retrieves the hello from world
// @Tags         hello
// @Produce      json
// @Success      200 {object}  ConnectedPeersResponse
// @Router       /peers [get]
func (s *Peer) peerDiscoveryHandler(w http.ResponseWriter, req *http.Request) {

	var response ConnectedPeersResponse
	response.Peers = make(map[string]string)

	for k, v := range s.p2p.sockets {
		response.Peers[k] = v.RemoteAddr().String()
	}
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode response as JSON and write to response writer
	json.
		NewEncoder(w).
		Encode(response)
}

type ConnectedPeersResponse struct {
	Peers map[string]string `json:"peers"`
}
