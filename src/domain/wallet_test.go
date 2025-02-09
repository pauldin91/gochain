package domain

import (
	"encoding/json"
	"log"
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

var tp = TransactionPool{}
var bc = Create()
var senderWallet = NewWallet(55.0)
var recipientWallet = NewWallet(44.0)

var testAmounts = []struct {
	amount           float64
	shouldBeExecuted bool
}{
	{5.0, true},
	{11.0, true},
	{22.0, true},
	{-19.0, false},
	{50000.0, false},
}

func TestCreateTransaction(t *testing.T) {

	executed := senderWallet.CreateTransaction(recipientWallet.address, testAmounts[0].amount, bc, &tp)
	if len(tp.transactions) != 1 || !executed {
		t.Errorf("should have %d but have %d\n", 1, len(tp.transactions))
	}
	tp.Clear()
}

func TestBalance(t *testing.T) {

	var senderBalance float64 = senderWallet.balance
	var recipientBalance float64 = recipientWallet.balance
	for _, ta := range testAmounts {
		executed := senderWallet.CreateTransaction(recipientWallet.address, ta.amount, bc, &tp)
		if executed != ta.shouldBeExecuted {
			t.Errorf("test with amount %0.8f it was supposed to %v while got %v", ta.amount, ta.shouldBeExecuted, executed)
			continue
		} else if !executed {
			continue
		}
		var tpDtos []TransactionData
		for _, t := range tp.transactions {
			tpDtos = append(tpDtos, t.Map())
		}
		jsonTransactions, _ := json.Marshal(tpDtos)

		bc.AddBlock(string(jsonTransactions))

		senderBalance = senderWallet.balance
		senderWallet.balance = senderWallet.CalculateBalance(bc)
		if senderWallet.balance != senderBalance-ta.amount {
			t.Errorf("Sender should have a balance of %0.8f but has %0.8f\n", senderBalance-ta.amount, senderWallet.balance)
		}
		recipientBalance = recipientWallet.balance
		recipientWallet.balance = recipientWallet.CalculateBalance(bc)
		if recipientWallet.balance != recipientBalance+ta.amount {
			t.Errorf("Recipient should have a balance of %0.8f but has %0.8f\n", recipientBalance+ta.amount, recipientWallet.balance)
		}

	}
	tp.Clear()
}

func TestVerify(t *testing.T) {

}
