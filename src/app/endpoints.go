package app

const (
	mineEndpoint             string = "/mine"
	blockEndpoint            string = "/blocks"
	balanceEndpoint          string = "/balance"
	publickeyEndpoint        string = "/public-key"
	transactionsEndpoint     string = "/transactions"
	mineTransactionsEndpoint string = transactionsEndpoint + mineEndpoint
)
