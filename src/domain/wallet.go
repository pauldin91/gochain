package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pauldin91/gochain/src/internal"
)

type Wallet struct {
	balance float64 `json:"balance"`
	keyPair *internal.KeyPair
	address string `json:"address"`
}

func (w Wallet) String() string {
	jsonWallet, _ := json.Marshal(w)
	return string(jsonWallet)
}

func NewWallet(init float64) *Wallet {
	res := Wallet{
		balance: init,
		keyPair: internal.NewKeyPair(),
	}
	res.address = res.keyPair.GetPublicKey()
	return &res
}

func (w Wallet) ToString() string {
	return fmt.Sprintf("Wallet - \npublicKey\t: %s\nbalance: %.8f\n", w.keyPair.GetPublicKey(), w.balance)
}

func (w Wallet) CalculateBalance(chain Blockchain) float64 {
	var totalTransactions []Transaction
	balance := w.balance
	for _, b := range chain.Chain {
		var transactions []Transaction
		_ = json.Unmarshal([]byte(b.Data), &transactions)
		totalTransactions = append(totalTransactions, transactions...)
	}
	walletInputTs := internal.FilterBy(totalTransactions, w.keyPair.GetPublicKey(), findTransactionByAddress)

	var start time.Time
	if len(walletInputTs) > 0 {
		recentInputT := internal.Reduce(walletInputTs, maxByTimestamp)
		balance = recentInputT.Output[w.keyPair.GetPublicKey()].Amount
		start = recentInputT.Input.Timestamp
	}

	v := TimestampAddressFilter{
		timestamp: start,
		address:   w.keyPair.GetPublicKey(),
	}

	filtered := internal.FilterBy(totalTransactions, v, findByAddressAndTimestamp)
	filteredOutputs := make(map[string]Input)
	for _, i := range filtered {
		for _, o := range i.Output {
			if o.Address == w.keyPair.GetPublicKey() {
				filteredOutputs[o.Address] = o
			}
		}
	}
	for _, b := range filteredOutputs {
		balance += b.Amount
	}

	return balance
}

func (w *Wallet) CreateTransaction(recipient string, amount float64, blockchain Blockchain, pool *TransactionPool) bool {

	w.balance = w.CalculateBalance(blockchain)
	if amount > w.balance || amount <= 0.0 {
		return false
	} else {
		transaction := NewTransaction(w, recipient, amount)
		pool.AddOrUpdateById(*transaction)
		return true
	}

}

func Verify(transaction Transaction) bool {
	return internal.VerifySignature(transaction.Input.Address, []byte(transaction.String()), []byte(transaction.Input.Signature))
}
