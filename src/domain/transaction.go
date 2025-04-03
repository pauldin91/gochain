package domain

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pauldin91/gochain/src/utils"
)

const (
	MINING_REWARD float64 = 3000
)

type Transaction struct {
	Id     uuid.UUID        `json:"id"`
	Input  Input            `json:"input"`
	Output map[string]Input `json:"output"`
	Amount float64          `json:"amount"`
}

func (t Transaction) String() string {
	jsonT, _ := json.Marshal(t)
	return string(jsonT)
}
func transactionWithOutputs(senderWallet Wallet, outputs []Input, amount float64) Transaction {
	transaction := Transaction{
		Id: uuid.New(),
	}
	transaction.Output = make(map[string]Input)
	for _, o := range outputs {

		transaction.Output[o.Address] = o
	}
	transaction.Amount = amount
	transaction.Input.Address = senderWallet.keyPair.GetPublicKey()
	transaction.Input.Amount = senderWallet.balance
	transaction.Input.Timestamp = time.Now().UTC()
	transaction.sign(senderWallet)
	return transaction
}

func (t *Transaction) sign(wallet Wallet) {
	outs, _ := json.Marshal(&t.Output)
	t.Input.Signature = wallet.keyPair.Sign(utils.Hash(string(outs)))
}

func NewTransaction(senderWallet Wallet, recipient string, amount float64) *Transaction {
	if amount > senderWallet.balance || amount <= 0 {
		return nil
	}
	outputs := []Input{
		{Amount: senderWallet.balance - amount, Address: senderWallet.keyPair.GetPublicKey(), Timestamp: time.Now().UTC()},
		{Amount: amount, Address: recipient, Timestamp: time.Now().UTC()},
	}
	var created Transaction = transactionWithOutputs(senderWallet, outputs, amount)
	return &created
}

func (t *Transaction) Update(senderWallet Wallet, recipientAddress string, amount float64) {
	senderOutput := t.Output[senderWallet.Address]
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
	t.sign(senderWallet)
	t.Output[newlyAdded.Address] = newlyAdded
}

func Verify(transaction Transaction) bool {
	outs, _ := json.Marshal(transaction.Output)
	var tsString string = utils.Hash(string(outs))
	return utils.VerifySignature(transaction.Input.Address, []byte(tsString), []byte(transaction.Input.Signature))
}

func Reward(minerWallet *Wallet, blockchainWallet *Wallet) *Transaction {
	outputs := []Input{
		{Amount: MINING_REWARD, Address: minerWallet.Address, Timestamp: time.Now().UTC()},
	}
	tr := transactionWithOutputs(*blockchainWallet, outputs, MINING_REWARD)

	return &tr
}
