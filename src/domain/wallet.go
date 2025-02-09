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
		var transactionsData []TransactionData
		_ = json.Unmarshal([]byte(b.Data), &transactionsData)
		for _, l := range transactionsData {
			transactions = append(transactions, l.Map())
		}
		totalTransactions = append(totalTransactions, transactions...)
	}
	walletInputTs := internal.FilterBy(totalTransactions, w.keyPair.GetPublicKey(), findTransactionByAddress)

	var start time.Time
	if len(walletInputTs) > 0 {
		recentInputT := internal.Reduce(walletInputTs, maxByTimestamp)
		balance = internal.FindBy(recentInputT.Output, w.keyPair.GetPublicKey(), findInputByAddress).Amount
		start = recentInputT.Input.Timestamp
	}

	v := TimestampAddressFilter{
		timestamp: start,
		address:   w.keyPair.GetPublicKey(),
	}

	filtered := internal.FilterBy(totalTransactions, v, findByAddressAndTimestamp)

	for _, i := range filtered {
		if i.Input.Timestamp.UnixMilli() > start.UnixMilli() {
			for _, c := range i.Output {
				if c.Address == w.address {
					balance += c.Amount
					break
				}
			}
		}
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
	return internal.VerifySignature(transaction.Input.Address, []byte(transaction.OutputString()), []byte(transaction.Input.Signature))
}
