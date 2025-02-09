package domain

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/pauldin91/gochain/src/internal"
)

type Wallet struct {
	balance float64
	keyPair *internal.KeyPair
	address string
}

func (w Wallet) ToString() string {
	return fmt.Sprintf("Wallet - \npublicKey\t: %s\nbalance: %.8f\n", w.keyPair.GetPublicKey(), w.balance)
}

func (w *Wallet) CalculateBalance(chain Blockchain) float64 {
	var totalTransactions []Transaction
	var transactions []Transaction
	balance := w.balance
	for _, b := range chain.Chain {
		_ = json.Unmarshal([]byte(b.Data), &transactions)
		totalTransactions = append(totalTransactions, transactions...)
	}
	walletInputTs := internal.FilterBy(totalTransactions, w.keyPair.GetPublicKey(), findTransactionByAddress)

	var start time.Time
	if len(walletInputTs) > 0 {
		recentInputT := internal.Reduce(walletInputTs, maxByTimestamp)
		balance = internal.FindBy(recentInputT.Output, w.keyPair.GetPublicKey(), findInputByAddress).amount
		start = recentInputT.Input.timestamp
	}

	v := TimestampAddressFilter{
		timestamp: start,
		address:   w.keyPair.GetPublicKey(),
	}

	filtered := internal.FilterBy(transactions, v, findByAddressAndTimestamp)

	internal.ForEachAction(filtered, &balance, func(b *float64, t Transaction) {
		*b += t.Input.amount
	})

	return balance
}

func (w *Wallet) createTransaction(recipient string, amount float64, blockchain Blockchain, pool *TransactionPool) {
	w.balance = w.CalculateBalance(blockchain)

	if amount > w.balance {
		log.Printf("Amount : %0.8f exceeds current balance %0.8f", amount, w.balance)
	} else {
		transaction := NewTransaction(w, recipient, amount)
		pool.AddOrUpdateById(*transaction)
	}

}

func Verify(transaction Transaction) bool {
	return internal.VerifySignature(transaction.Input.address, []byte(transaction.OutputString()), []byte(transaction.Input.signature))
}
