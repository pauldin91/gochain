package domain

import (
	"time"

	"github.com/google/uuid"
)

type TransactionData struct {
	timestamp time.Time
	amount    float64
	address   string
	signature string
}

type Transaction struct {
	Id     uuid.UUID
	Input  TransactionData
	Output []map[string]TransactionData
}

func (t *Transaction) Update(senderAddress string, recipientAddress string, amount float64) {

}

func createTransaction(recipient string, amount float64, blockchain Blockchain, pool TransactionPool) {

}


