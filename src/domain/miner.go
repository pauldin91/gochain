package domain

import "encoding/json"

type Miner struct {
	Bc     *Blockchain
	Tp     *TransactionPool
	Wallet *Wallet
}

func (ner *Miner) Mine(blockchainWallet *Wallet) *Block {
	validTransactions := ner.Tp.ValidTransactions()
	validTransactions = append(validTransactions, *Reward(ner.Wallet, blockchainWallet))

	data, _ := json.Marshal(validTransactions)
	block := ner.Bc.AddBlock(string(data))
	return &block
}

func (ner *Miner) Clear() {
	ner.Tp.Clear()
}
