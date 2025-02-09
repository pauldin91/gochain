package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
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

func (t *Input) sign(wallet *Wallet) {
	t.Signature = wallet.keyPair.Sign(internal.Hash(t.String()))
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

func (t *TransactionData) Map() Transaction {

	var outputs []Input
	for _, t := range t.Output {
		var out Input
		_ = json.Unmarshal([]byte(t), &out)
		outputs = append(outputs, out)
	}
	var input Input
	_ = json.Unmarshal([]byte(t.Input), &input)
	return Transaction{
		Id:     t.Id,
		Input:  input,
		Output: outputs,
	}
}
