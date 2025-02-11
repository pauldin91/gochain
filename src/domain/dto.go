package domain

import (
	"encoding/json"
	"time"

	"github.com/pauldin91/gochain/src/internal"
)

type Input struct {
	Timestamp time.Time `json:"timestamp"`
	Amount    float64   `json:"amount"`
	Address   string    `json:"address"`
	Signature string    `json:"signature"`
}

func (i Input) GetAmount() float64 {
	return i.Amount
}

func (i Input) GetAddress() string {
	return i.Address
}

func (t Input) String() string {
	res, _ := json.Marshal(t)
	return string(res)
}

func (t *Input) sign(wallet Wallet) {
	t.Signature = wallet.keyPair.Sign(internal.Hash(t.String()))
}

type TimestampAddressFilter struct {
	timestamp time.Time
	address   string
}
