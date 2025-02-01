package app

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/pauldin91/gochain/src/internal"
)

type Server interface {
	WithConfig()
}

type HttpServer struct {
	cfg    internal.Config
	router *chi.Mux
	p2p    WsServer
}

type ServerBuilder struct {
	Server *HttpServer
}

func (sb ServerBuilder) WithConfig(settings string) ServerBuilder {
	cfg, err := internal.LoadConfig(settings)
	if err != nil {
		log.Fatal("unable to load config")
	}
	sb.Server.cfg = cfg
	return sb
}

func (sb ServerBuilder) WithRouter() ServerBuilder {
	sb.Server.router = chi.NewRouter()
	sb.Server.router.Get(balanceEndpoint, balanceHandler)
	sb.Server.router.Get(blockEndpoint, blockHandler)
	sb.Server.router.Get(mineEndpoint, mineHandler)
	sb.Server.router.Get(transactionsEndpoint, getTransactionsHandler)
	sb.Server.router.Post(transactionsEndpoint, createTransactionHandler)
	return sb
}

func (sb ServerBuilder) Build() *HttpServer {
	return sb.Server
}

func (s *HttpServer) Start() {

	certFile := filepath.Join(s.cfg.CertPath, s.cfg.CertFile)
	certKey := filepath.Join(s.cfg.CertPath, s.cfg.CertKey)
	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		log.Fatal("unable to load certs")
	}

	server := &http.Server{
		Addr:    s.cfg.HttpServerAddress,
		Handler: s.router,
	}

	err := server.ListenAndServeTLS(certFile, certKey)
	if err != nil {
		log.Fatal("Could not start server")
	}
}


