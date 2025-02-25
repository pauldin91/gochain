package app

import (
	"log"

	"github.com/go-chi/chi/v5"
	_ "github.com/pauldin91/gochain/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/pauldin91/gochain/src/internal"
)

type ServerBuilder struct {
	Server *HttpServer
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{
		Server: &HttpServer{},
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

func (sb *ServerBuilder) WithRouter() *ServerBuilder {
	sb.Server.router = chi.NewRouter()
	sb.Server.router.Get("/swagger/*", httpSwagger.WrapHandler)
	sb.Server.router.Get(balanceEndpoint, balanceHandler)
	sb.Server.router.Post(blockEndpoint, blockHandler)
	sb.Server.router.Get(mineEndpoint, mineHandler)
	sb.Server.router.Get(publickeyEndpoint, publicKeyHandler)
	sb.Server.router.Get(transactionsEndpoint, getTransactionsHandler)
	sb.Server.router.Post(transactionsEndpoint, createTransactionHandler)
	return sb
}

func (sb *ServerBuilder) Build() *HttpServer {
	return sb.Server
}
