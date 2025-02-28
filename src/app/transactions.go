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
func (s *HttpServer) getTransactionsHandler(writer http.ResponseWriter, req *http.Request) {
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
// @Success      200
// @Router       /transactions [post]
func (s *HttpServer) createTransactionHandler(writer http.ResponseWriter, req *http.Request) {
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
