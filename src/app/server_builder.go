package app

import (
	"log"

	"github.com/go-chi/chi/v5"
	_ "github.com/pauldin91/gochain/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/pauldin91/gochain/src/domain"
	"github.com/pauldin91/gochain/src/internal"
)

type ServerBuilder struct {
	Server *HttpServer
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{
		Server: &HttpServer{
			chain: domain.Create(),
		},
	}
}

func (sb *ServerBuilder) WithConfig(settings string) *ServerBuilder {
	cfg, err := internal.LoadConfig(settings)
	if err != nil {
		log.Fatal("unable to load config")
	}
	sb.Server.cfg = cfg
	return sb
}

func (sb *ServerBuilder) WithWsServer() *ServerBuilder {
	sb.Server.p2p = &WsServer{}
	return sb
}

func (sb *ServerBuilder) WithPool() *ServerBuilder {
	sb.Server.pool = &domain.TransactionPool{}
	return sb
}

func (sb *ServerBuilder) WithWallet() *ServerBuilder {
	wallet := domain.NewWallet(0.0)
	sb.Server.wallet = &wallet
	return sb
}

func (sb *ServerBuilder) WithRouter() *ServerBuilder {
	sb.Server.router = chi.NewRouter()
	sb.Server.router.Get("/swagger/*", httpSwagger.WrapHandler)
	sb.Server.router.Get(balanceEndpoint, balanceHandler)
	sb.Server.router.Get(blockEndpoint, sb.Server.blockHandler)
	sb.Server.router.Post(mineEndpoint, sb.Server.mineHandler)
	sb.Server.router.Get(publickeyEndpoint, sb.Server.publicKeyHandler)
	sb.Server.router.Get(transactionsEndpoint, sb.Server.getTransactionsHandler)
	sb.Server.router.Post(transactionsEndpoint, sb.Server.createTransactionHandler)
	return sb
}

func (sb *ServerBuilder) Build() *HttpServer {
	return sb.Server
}
