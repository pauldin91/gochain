package app

import (
	"net/http"

	_ "github.com/pauldin91/gochain/src/domain"
)

// retrieves transactions
// @Summary      Gets transactions
// @Description  Gets transactions
// @Tags         transactions
// @Produce      json
// @Success      200  {array}   domain.Transaction
// @Router       /transactions [get]
func getTransactionsHandler(writer http.ResponseWriter, req *http.Request) {

}

// creates a transaction
// @Summary      Create transaction
// @Description  Creates a transaction in the pool
// @Tags         transactions
// @Produce      json
// @Success      200  {array}   domain.Transaction
// @Router       /transactions [post]
func createTransactionHandler(writer http.ResponseWriter, req *http.Request) {
}
