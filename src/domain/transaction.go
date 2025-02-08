package domain

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pauldin91/gochain/src/internal"
)

type TransactionData struct {
	timestamp time.Time `json:"timestamp"`
	amount    float64   `json:"amount"`
	address   string    `json:"address"`
	signature string    `json:"signature"`
}

func (t *TransactionData) String() string {
	return fmt.Sprintf("%s | %0.8f | %s", t.address, t.amount, t.timestamp.Format(time.RFC3339))
}

type Transaction struct {
	Id     uuid.UUID                  `json:"id"`
	Input  TransactionData            `json:"input"`
	Output map[string]TransactionData `json:"output"`
}

func (t *Transaction) OutputString() string {
	var outputs []string
	for _, v := range t.Output {
		outputs = append(outputs, v.String())
	}
	return strings.Join(outputs, "|")
}

func NewTransaction(w Wallet, recipient string, amount float64) Transaction {
	return Transaction{}
}

func (t *Transaction) Update(senderWallet Wallet, recipientAddress string, amount float64) {
	senderOutput := t.Output[senderWallet.address]

	if amount > senderOutput.amount {
		log.Printf("amount %0.8f exceeds balance %0.8f", amount, senderWallet.balance)
		return
	}
	senderOutput.amount = senderOutput.amount - amount
	newlyAdded := TransactionData{
		timestamp: time.Now().UTC(),
		amount:    amount,
		address:   recipientAddress,
	}
	newlyAdded.sign(senderWallet)

	t.Output[senderWallet.address] = newlyAdded

}

func (t *TransactionData) sign(wallet Wallet) {
	t.signature = wallet.keyPair.Sign(internal.Hash(t.String()))
}
