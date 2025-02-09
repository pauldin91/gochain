package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pauldin91/gochain/src/internal"
)

type Input struct {
	timestamp time.Time `json:"timestamp"`
	amount    float64   `json:"amount"`
	address   string    `json:"address"`
	signature string    `json:"signature"`
}

func (i Input) GetAmount() float64 {
	return i.amount
}

func (i Input) GetAddress() string {
	return i.address
}

func (t Input) String() string {
	return fmt.Sprintf("%s | %0.8f | %s", t.address, t.amount, t.timestamp.Format(time.RFC3339))
}

func (t *Input) sign(wallet *Wallet) {
	t.signature = wallet.keyPair.Sign(internal.Hash(t.String()))
}

type TimestampAddressFilter struct {
	timestamp time.Time
	address   string
}

type TransactionData struct {
	Id     uuid.UUID `json:"id"`
	Input  string    `json:"input"`
	Output []string  `json:"outputs"`
}


