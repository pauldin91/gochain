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

func NewWallet(init float64) Wallet {
	res := Wallet{
		balance: init,
		keyPair: internal.NewKeyPair(),
	}
	res.address = res.keyPair.GetPublicKey()
	return res
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

	filteredOutputs := make(map[string]Input)
	internal.Flattened(totalTransactions, &filteredOutputs, func(t *Transaction, m *map[string]Input) {
		for _, i := range t.Output {
			if i.Address == v.address && t.Input.Timestamp.After(v.timestamp) {
				filteredOutputs[i.Address] = i
			}
		}
	})

	for _, b := range filteredOutputs {
		balance += b.Amount
	}

	return balance
}

func (w Wallet) CreateTransaction(recipient string, amount float64, blockchain Blockchain, pool *TransactionPool) bool {

	w.balance = w.CalculateBalance(blockchain)
	if amount > w.balance || amount <= 0.0 {
		return false
	} else {
		transaction := NewTransaction(w, recipient, amount)
		pool.AddOrUpdateById(transaction)
		return true
	}

}


