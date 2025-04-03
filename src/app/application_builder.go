package app

import (
	"log"

	"github.com/go-chi/chi/v5"
	_ "github.com/pauldin91/gochain/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/pauldin91/gochain/src/internal"
)

type WebApplicationBuilder struct {
	app  *HttpApplication
	peer *Peer
}

func NewServerBuilder(peer *Peer) *WebApplicationBuilder {

	return &WebApplicationBuilder{
		app:  &HttpApplication{},
		peer: peer,
	}
}

func (sb *WebApplicationBuilder) WithConfig(settings string) *WebApplicationBuilder {
	cfg, err := internal.LoadConfig(settings)
	if err != nil {
		log.Fatal("unable to load config")
	}
	sb.app.cfg = cfg
	return sb
}

func (sb *WebApplicationBuilder) WithRouter() *WebApplicationBuilder {
	sb.app.router = chi.NewRouter()

	return sb
}

func (sb *WebApplicationBuilder) Build() *HttpApplication {
	return sb.app.
		AddGet(blockEndpoint, sb.peer.blockHandler).
		AddGet(peerDiscoveryEndpoint, sb.peer.peerDiscoveryHandler).
		AddGet(balanceEndpoint, sb.peer.balanceHandler).
		AddGet(publickeyEndpoint, sb.peer.publicKeyHandler).
		AddGet(swaggerDocsEndpoint, httpSwagger.WrapHandler).
		AddGet(transactionsEndpoint, sb.peer.getTransactionsHandler).
		AddPost(mineBlockEndpoint, sb.peer.mineBlockHandler).
		AddPost(transactionsEndpoint, sb.peer.createTransactionHandler).
		AddPost(mineTransactionsEndpoint, sb.peer.mineTransactionHandler)
}
