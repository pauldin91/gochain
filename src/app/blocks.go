package app

import (
	"encoding/json"
	"net/http"
)

// gets the blockchain
// @Summary      Gets the blockchain
// @Description  Gets the blockchain
// @Tags         blocks
// @Produce      json
// @Success      200 {object} BlockchainDto
// @Router       /blocks [get]
func (s *HttpApplication) blockHandler(writer http.ResponseWriter, req *http.Request) {
	chain := BlockchainDto{}
	chain.Map(*s.peer.chain)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(chain)
}

// mine a block
// @Summary      Mine a block
// @Description  Mines a new block and adds it to the blockchain
// @Tags         mine
// @Accept       json
// @Produce      json
// @Param        request body BlockRequestDto true "Block Request Data"
// @Success      200 {object} BlockResponseDto
// @Failure      400 {object} map[string]string "Invalid request data"
// @Router       /mine [post]
func (s *HttpApplication) mineBlockHandler(writer http.ResponseWriter, req *http.Request) {
	block := BlockRequestDto{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&block)
	if err != nil {
		writeErrorResponse(writer, 400, "Invalid data")
		return
	}
	addedBlock := s.peer.chain.AddBlock(block.Data)
	dto := &BlockResponseDto{}
	dto.Map(addedBlock)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(dto)
}
