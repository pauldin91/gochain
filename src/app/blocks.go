package app

import (
	"encoding/json"
	"net/http"

	"github.com/pauldin91/gochain/src/domain"
)

// blockHandler creates a new block in the chain
// @Summary      Creates a block
// @Description  Creats a block
// @Tags         block
// @Produce      json
// @Success      200  {array}   Block
// @Router       /blocks [post]

func (ws HttpServer) blockHandler(w http.ResponseWriter, req *http.Request) {
	block := domain.Block{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}
	chain.AddBlock(block.ToString())

	w.WriteHeader(http.StatusCreated)
}
