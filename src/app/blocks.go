package app

import (
	"encoding/json"
	"net/http"

	"github.com/pauldin91/gochain/src/domain"
)

func (ws HttpServer) blockHandler(w http.ResponseWriter, req *http.Request) {
	block := domain.Block{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}
	chain.AddBlock(block.ToString(), int64(ws.cfg.MineRate))

	w.WriteHeader(http.StatusCreated)
}
