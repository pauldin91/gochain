package domain

import "testing"

var amount float64 = 10.0
var transaction = NewTransaction(senderWallet, recipientWallet.Address, amount)

func TestNewTransaction(t *testing.T) {

	if len(transaction.Output) != 2 {
		t.Error("Transaction output length is always 2")
	}

	sender, ex := transaction.Output[senderWallet.Address]
	if !ex ||
		sender.Amount != senderWallet.balance-amount ||
		sender.Address != senderWallet.Address {
		t.Error("Invalid sender transaction")
	}

	recipient, ex := transaction.Output[recipientWallet.Address]
	if !ex ||
		recipient.Amount != recipientWallet.balance+amount ||
		recipient.Address != recipientWallet.Address {
		t.Error("Invalid recipient transaction")
	}
}

func TestVerifyTransaction(t *testing.T) {
	res := Verify(*transaction)
	if !res {
		t.Error("Valid transaction should be validated")
	}
	var copy Input = transaction.Output[recipientWallet.Address]

	copy.Amount = 30000
	transaction.Output[recipientWallet.Address] = copy

	res = Verify(*transaction)
	if res {
		t.Error("invalid transaction should be invalidated")
	}

}
