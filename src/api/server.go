package api

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/pauldin91/gochain/src/internal"
)

type Server struct {
	cfg    internal.Config
	router *chi.Mux
}

func (s *Server) Start() {

	certFile := filepath.Join(s.cfg.CertPath, s.cfg.CertFile)
	certKey := filepath.Join(s.cfg.CertPath, s.cfg.CertKey)

	err := http.ListenAndServeTLS(":"+string(s.cfg.HttpServerAddress), certFile, certKey, s.router)
	if err != nil {
		log.Fatal("Could not start server")
	}
}

type ServerBuilder struct {
	Server *Server
}

func (sb ServerBuilder) WithConfig(cfg internal.Config) ServerBuilder {
	sb.Server.cfg = cfg
	return sb
}

func (sb ServerBuilder) WithRouter() ServerBuilder {
	sb.Server.router = chi.NewRouter()
	sb.Server.router.Get(blockEndpoint, blockHandler)
	sb.Server.router.Get(mineEndpoint, mineHandler)
	sb.Server.router.Get(transactionsEndpoint, getTransactionsHandler)
	sb.Server.router.Post(transactionsEndpoint, createTransactionHandler)
	return sb
}

func (sb ServerBuilder) Build() *Server {
	return sb.Server
}
