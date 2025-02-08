package domain

type TransactionPool struct {
	transactions map[string]Transaction
}

func (p *TransactionPool) AddOrUpdateById(transaction Transaction) {
	p.transactions[transaction.Id.String()] = transaction

}

func (p *TransactionPool) TransactionByAddress(address string) *Transaction {
	for _, t := range p.transactions {
		if t.Input.address == address {
			return &t
		}
	}
	return nil
}

func (p *TransactionPool) ValidTransactions() []Transaction {
	var transactions []Transaction
	for _, t := range p.transactions {
		transaction := filter(t)
		if transaction != nil {
			transactions = append(transactions, *transaction)
		}
	}
	return transactions
}
func (p *TransactionPool) Clear() {
	p.transactions = make(map[string]Transaction)
}

func filter(transaction Transaction) *Transaction {
	var totalOutput float64 = 0.0
	for _, z := range transaction.Output {
		totalOutput += z.amount
	}
	if transaction.Input.amount != totalOutput {
		return nil
	}

	return &transaction

}
