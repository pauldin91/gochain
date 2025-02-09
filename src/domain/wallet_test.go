package domain

import (
	"encoding/json"
	"testing"

	"github.com/pauldin91/gochain/src/internal"
)

func blockchainWallet() Wallet {
	return Wallet{
		balance: 30.0,
		keyPair: internal.NewKeyPair(),
		address: "hakuna-mattata",
	}
}

var wallet = blockchainWallet()

var tp = TransactionPool{}

var bc = Create()

func TestCreateTransaction(t *testing.T) {

	var amount float64 = 5.0
	recipient := "mattata-hakuna"
	wallet.createTransaction(recipient, amount, bc, &tp)
	if len(tp.transactions) != 1 {
		t.Errorf("should have %d but have %d\n", 2, len(tp.transactions))
	}
	tp.Clear()
}

func TestBalance(t *testing.T) {
	var amount float64 = 5.0
	recipient := "mattata-hakuna"
	wallet.createTransaction(recipient, amount, bc, &tp)
	var tpDtos []TransactionData
	for _, t := range tp.transactions {
		tpDtos = append(tpDtos, t.Map())
	}

	jsonTransactions, _ := json.Marshal(tpDtos)

	bc.AddBlock(string(jsonTransactions))
	wallet.balance = wallet.CalculateBalance(bc)
	if wallet.balance != 25.0 {
		t.Errorf("should have %0.8f but have %0.8f\n", 25.0, wallet.balance)
	}
}

func TestVerify(t *testing.T) {

}
