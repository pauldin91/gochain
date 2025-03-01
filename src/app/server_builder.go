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
	peer   *Peer
}

func NewServerBuilder(peer *Peer) *ServerBuilder {

	return &ServerBuilder{
		Server: &HttpServer{},
		peer:   peer,
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

func (sb *ServerBuilder) WithRouter() *ServerBuilder {
	sb.Server.router = chi.NewRouter()

	return sb
}

func (sb *ServerBuilder) Build() *HttpServer {
	return sb.Server.
		AddGet(blockEndpoint, sb.peer.blockHandler).
		AddGet(balanceEndpoint, sb.peer.balanceHandler).
		AddGet(publickeyEndpoint, sb.peer.publicKeyHandler).
		AddGet(swaggerDocsEndpoint, httpSwagger.WrapHandler).
		AddGet(transactionsEndpoint, sb.peer.getTransactionsHandler).
		AddPost(mineBlockEndpoint, sb.peer.mineBlockHandler).
		AddPost(transactionsEndpoint, sb.peer.createTransactionHandler).
		AddPost(mineTransactionsEndpoint, sb.peer.mineTransactionHandler)
}
