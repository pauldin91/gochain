package domain

import "testing"

var amount float64 = 10.0
var transaction = NewTransaction(senderWallet, recipientWallet.address, amount)

func TestNewTransaction(t *testing.T) {

	if len(transaction.Output) != 2 {
		t.Error("Transaction output length is always 2")
	}

	sender, ex := transaction.Output[senderWallet.address]
	if !ex ||
		sender.Amount != senderWallet.balance-amount ||
		sender.Address != senderWallet.address {
		t.Error("Invalid sender transaction")
	}

	recipient, ex := transaction.Output[recipientWallet.address]
	if !ex ||
		recipient.Amount != recipientWallet.balance+amount ||
		recipient.Address != recipientWallet.address {
		t.Error("Invalid recipient transaction")
	}
}

func TestVerifyTransaction(t *testing.T) {
	res := Verify(transaction)
	if !res {
		t.Error("Valid transaction should be validated")
	}
	var copy Input = transaction.Output[recipientWallet.address]

	copy.Amount = 30000
	transaction.Output[recipientWallet.address] = copy

	res = Verify(transaction)
	if res {
		t.Error("invalid transaction should be invalidated")
	}

}
