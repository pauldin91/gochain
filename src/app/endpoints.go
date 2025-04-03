package app

const (
	mineBlockEndpoint        string = "/mine"
	peerDiscoveryEndpoint    string = "/peers"
	blockEndpoint            string = "/blocks"
	balanceEndpoint          string = "/balance"
	publickeyEndpoint        string = "/public-key"
	transactionsEndpoint     string = "/transactions"
	mineTransactionsEndpoint string = transactionsEndpoint + mineBlockEndpoint
	swaggerDocsEndpoint      string = "/swagger/*"
)
