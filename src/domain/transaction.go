package domain

import (
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pauldin91/gochain/src/internal"
)

type Transaction struct {
	Id     uuid.UUID `json:"id"`
	Input  Input     `json:"input"`
	Output []Input   `json:"output"`
}

func (t Transaction) OutputString() string {
	var outputs []string
	for _, v := range t.Output {
		outputs = append(outputs, v.String())
	}
	return strings.Join(outputs, "|")
}
func transactionWithOutputs(senderWallet *Wallet, outputs []Input) Transaction {
	transaction := Transaction{
		Id: uuid.New(),
	}
	transaction.Output = append(transaction.Output, outputs...)
	transaction.Input.Address = senderWallet.keyPair.GetPublicKey()
	transaction.Input.Amount = senderWallet.balance
	transaction.Input.Timestamp = time.Now().UTC()
	transaction.Input.sign(senderWallet)
	return transaction
}

func NewTransaction(senderWallet *Wallet, recipient string, amount float64) *Transaction {
	if amount > senderWallet.balance {
		//log.Errorf("Amount : %0.8f exceeds balance %0.8f\n", amount, w.balance)
		return nil
	}
	outputs := []Input{
		{Amount: senderWallet.balance - amount, Address: senderWallet.keyPair.GetPublicKey(), Timestamp: time.Now().UTC()},
		{Amount: amount, Address: recipient, Timestamp: time.Now().UTC()},
	}
	var created Transaction = transactionWithOutputs(senderWallet, outputs)
	return &created
}

func (t *Transaction) Update(senderWallet *Wallet, recipientAddress string, amount float64) {
	senderOutput := *internal.FindBy(t.Output, senderWallet.address, findInputByAddress)
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
	t.Output = append(t.Output, newlyAdded)
}

func (t *Transaction) Map() TransactionData {

	var outputs []string
	for _, t := range t.Output {
		outputs = append(outputs, t.String())
	}
	return TransactionData{
		Id:     t.Id,
		Input:  t.Input.String(),
		Output: outputs,
	}
}
