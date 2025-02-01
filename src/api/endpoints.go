package api

const (
	mineEndpoint             string = "/mine"
	blockEndpoint            string = "/blocks"
	publickeyEndpoint        string = "/public-key"
	transactionsEndpoint     string = "/transactions"
	mineTransactionsEndpoint string = transactionsEndpoint + mineEndpoint
)
