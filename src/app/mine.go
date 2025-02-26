package app

import (
	"encoding/json"
	"net/http"
)

type BlockDto struct {
	Data string `json:"data"`
}

// mines a transaction
// @Summary      Mine transaction
// @Description  Mines transaction
// @Tags         mine
// @Produce      json
// @Success      200 {object} BlockDto
// @Router       /mine [get]
func mineHandler(writer http.ResponseWriter, req *http.Request) {
	block := BlockDto{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&block)
	if err != nil {
		writeErrorResponse(writer, 400, "Invalid data")
		return
	}
	chain.AddBlock(block.Data)
	writer.WriteHeader(http.StatusCreated)
}
