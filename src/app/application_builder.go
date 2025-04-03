package app

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	_ "github.com/pauldin91/gochain/docs"
	"github.com/pauldin91/gochain/src/utils"
	httpSwagger "github.com/swaggo/http-swagger"
)

type WebApplicationBuilder struct {
	app *HttpApplication
}

func NewServerBuilder() *WebApplicationBuilder {

	return &WebApplicationBuilder{
		app: &HttpApplication{},
	}
}

func (sb *WebApplicationBuilder) WithConfig(settings string) *WebApplicationBuilder {
	cfg, err := utils.LoadConfig(settings)
	if err != nil {
		log.Fatal("unable to load config")
	}
	sb.app.cfg = cfg
	return sb
}

func (sb *WebApplicationBuilder) WithPeerServer() *WebApplicationBuilder {
	sb.app.ws = &WsServer{
		cfg: sb.app.cfg,
	}
	sb.app.ws.sockets = make(map[string]*websocket.Conn)
	if sb.app.cfg.Peers != "" {
		go sb.app.ws.connectToPeers(sb.app.cfg.Peers)
	}

	return sb
}

func (sb *WebApplicationBuilder) WithRouter() *WebApplicationBuilder {
	sb.app.router = chi.NewRouter()

	return sb
}

func (sb *WebApplicationBuilder) Build() *HttpApplication {
	return sb.app.
		AddGet(blockEndpoint, sb.app.blockHandler).
		AddGet(peerDiscoveryEndpoint, sb.app.peerDiscoveryHandler).
		AddGet(balanceEndpoint, sb.app.balanceHandler).
		AddGet(publickeyEndpoint, sb.app.publicKeyHandler).
		AddGet(swaggerDocsEndpoint, httpSwagger.WrapHandler).
		AddGet(transactionsEndpoint, sb.app.getTransactionsHandler).
		AddPost(mineBlockEndpoint, sb.app.mineBlockHandler).
		AddPost(transactionsEndpoint, sb.app.createTransactionHandler).
		AddPost(mineTransactionsEndpoint, sb.app.mineTransactionHandler)
}
