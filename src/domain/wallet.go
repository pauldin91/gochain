package domain

import (
	"fmt"
	"log"

	"github.com/pauldin91/gochain/src/internal"
)

type Wallet struct {
	balance float64
	keyPair internal.KeyPair
	address string
}

func (w Wallet) ToString() string {
	return fmt.Sprintf("Wallet - \npublicKey\t: %s\nbalance: %.8f\n", w.keyPair.GetPublicKey(), w.balance)
}

func (w *Wallet) CalculateBalance(chain Blockchain) float64 {
	//var transactions []Transaction
	balance := w.balance

	return balance
}

func (w Wallet) createTransaction(recipient string, amount float64, blockchain Blockchain, pool TransactionPool) {
	w.balance = w.CalculateBalance(blockchain)

	if amount > w.balance {
		log.Printf("Amount : %0.8f exceeds current balance %0.8f", amount, w.balance)
	} else {
		transaction := NewTransaction(w, recipient, amount)
		pool.AddOrUpdateById(transaction)
	}

}

func Verify(transaction Transaction) bool {
	return internal.VerifySignature(transaction.Input.address, []byte(transaction.OutputString()), []byte(transaction.Input.signature))
}

func blockchainWallet() Wallet {
	return Wallet{
		balance: 0.0,
		keyPair: internal.NewKeyPair(),
		address: "hakuna-mattata",
	}
}
