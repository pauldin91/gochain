package app

import (
	"encoding/json"
	"net/http"

	_ "github.com/pauldin91/gochain/src/domain"
)

// retrieves transactions
// @Summary      Gets transactions
// @Description  Gets transactions
// @Tags         transactions
// @Produce      json
// @Success      200  {object} PoolDto
// @Router       /transactions [get]
func (s *Peer) getTransactionsHandler(writer http.ResponseWriter, req *http.Request) {
	tp := PoolDto{}
	tp.Map(s.pool)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tp)
}

// creates a transaction
// @Summary      Create transaction
// @Description  Creates a transaction in the pool
// @Tags         transactions
// @Produce      json
// @Param        request body TransactionRequestDto true "Transaction Request Data"
// @Success      200
// @Router       /transactions [post]
func (s *Peer) createTransactionHandler(writer http.ResponseWriter, req *http.Request) {
	t := TransactionRequestDto{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&t)
	if err != nil {
		writeErrorResponse(writer, 400, "Invalid data")
		return
	}
	s.wallet.CreateTransaction(t.Recipient, t.Amount, *s.chain, s.pool)
	writer.WriteHeader(http.StatusCreated)
}

// mines a transaction
// @Summary      mines transaction
// @Description  mines a transaction in the pool
// @Tags         transactions
// @Produce      json
// @Success      200
// @Router       /transactions/mine [post]
func (s *Peer) mineTransactionHandler(writer http.ResponseWriter, req *http.Request) {
	block := s.mine()
	dto := BlockResponseDto{}
	dto.Map(block)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(dto)

}
