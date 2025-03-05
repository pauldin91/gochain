package domain

type TransactionPool struct {
	Transactions []Transaction
}

func (p *TransactionPool) Size() int {
	return len(p.Transactions)
}

func (p *TransactionPool) AddOrUpdateById(transaction Transaction) {
	var t *Transaction = nil
	for i, tr := range p.Transactions {
		if tr.Id == transaction.Id {
			p.Transactions[i] = transaction
			break
		}
	}
	if t == nil && transaction.Amount > 0 {
		p.Transactions = append(p.Transactions, transaction)
	}

}

func (p *TransactionPool) TransactionById(id string) *Transaction {
	for _, t := range p.Transactions {
		if t.Id.String() == id {
			return &t
		}
	}
	return nil
}

func (p *TransactionPool) ValidTransactions() []Transaction {
	var transactions []Transaction
	for _, t := range p.Transactions {
		transaction := filter(t)
		if transaction != nil {
			transactions = append(transactions, *transaction)
		}
	}
	return transactions
}
func (p *TransactionPool) Clear() {
	p.Transactions = []Transaction{}
}

func filter(transaction Transaction) *Transaction {
	var totalOutput float64 = 0.0
	for _, z := range transaction.Output {
		totalOutput += z.Amount
	}
	if transaction.Input.Amount != totalOutput {
		return nil
	}
	if !Verify(transaction) {
		return nil
	}

	return &transaction

}
