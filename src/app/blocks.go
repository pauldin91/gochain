package app

import (
	"encoding/json"
	"net/http"

	"github.com/pauldin91/gochain/src/domain"
)

type BlockchainDto struct {
	Chain string `json:"chain"`
}

func (dto *BlockchainDto) Map(ch domain.Blockchain) {
	dto.Chain = ch.String()
}

// creates a new block in the chain
// @Summary      Creates a block
// @Description  Creates a block
// @Tags         blocks
// @Produce      json
// @Success      200 {object} BlockchainDto
// @Router       /blocks [get]
func (s *HttpServer) blockHandler(w http.ResponseWriter, req *http.Request) {
	chain := BlockchainDto{}
	chain.Map(*s.chain)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chain)
}
