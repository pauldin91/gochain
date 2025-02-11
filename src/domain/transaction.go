package domain

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id     uuid.UUID        `json:"id"`
	Input  Input            `json:"input"`
	Output map[string]Input `json:"output"`
}

func (t Transaction) String() string {
	jsonT, _ := json.Marshal(t)
	return string(jsonT)
}
func transactionWithOutputs(senderWallet Wallet, outputs []Input) Transaction {
	transaction := Transaction{
		Id: uuid.New(),
	}
	transaction.Output = make(map[string]Input)
	for _, o := range outputs {

		transaction.Output[o.Address] = o
	}
	transaction.Input.Address = senderWallet.keyPair.GetPublicKey()
	transaction.Input.Amount = senderWallet.balance
	transaction.Input.Timestamp = time.Now().UTC()
	transaction.Input.sign(senderWallet)
	return transaction
}

func NewTransaction(senderWallet Wallet, recipient string, amount float64) *Transaction {
	if amount > senderWallet.balance {
		return nil
	}
	outputs := []Input{
		{Amount: senderWallet.balance - amount, Address: senderWallet.keyPair.GetPublicKey(), Timestamp: time.Now().UTC()},
		{Amount: amount, Address: recipient, Timestamp: time.Now().UTC()},
	}
	var created Transaction = transactionWithOutputs(senderWallet, outputs)
	return &created
}

func (t *Transaction) Update(senderWallet Wallet, recipientAddress string, amount float64) {
	senderOutput := t.Output[senderWallet.address]
	if amount > senderOutput.Amount {
		log.Printf("amount %0.8f exceeds balance %0.8f", amount, senderWallet.balance)
		return
	}
	senderOutput.Amount = senderOutput.Amount - amount
	newlyAdded := Input{
		Timestamp: time.Now().UTC(),
		Amount:    amount,
		Address:   recipientAddress,
	}
	newlyAdded.sign(senderWallet)
	t.Output[newlyAdded.Address] = newlyAdded
}
