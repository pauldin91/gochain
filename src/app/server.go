package app

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/pauldin91/gochain/src/internal"
)

type Server interface {
	WithConfig()
}

type HttpServer struct {
	cfg    internal.Config
	router *chi.Mux
	p2p    *WsServer
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

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  int(s.cfg.WsReadLimit),
			WriteBufferSize: int(s.cfg.WsWriteLimit),
		}
		ws, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Printf("could not establish a connection with %s", ws.RemoteAddr())
		}

		ws.SetReadLimit(s.cfg.WsReadLimit)
		s.p2p.sockets = append(s.p2p.sockets, ws)
		s.p2p.broadcast(chain.Chain)
	})

	err := server.ListenAndServeTLS(certFile, certKey)
	if err != nil {
		log.Fatal("Could not start server")
	}
	log.Printf("INFO : http server starter on %s\n", s.cfg.HttpServerAddress)
	err = http.ListenAndServeTLS(s.cfg.WsServerAddress, certFile, certKey, nil)
	if err != nil {
		log.Fatal("Could not start ws server", err)
	}
	log.Printf("INFO : ws server starter on %s\n", s.cfg.WsServerAddress)
}

