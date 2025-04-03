package app

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/pauldin91/gochain/src/internal"
)

type Application interface {
	Start()
}

type HttpApplication struct {
	cfg    internal.Config
	router *chi.Mux
}

func (s *HttpApplication) Start(peer *Peer) {
	certFile := filepath.Join(s.cfg.CertPath, s.cfg.CertFile)
	certKey := filepath.Join(s.cfg.CertPath, s.cfg.CertKey)

	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		log.Fatal("unable to load certs")
	}

	server := &http.Server{
		Addr:    s.cfg.HttpServerAddress,
		Handler: s.router,
	}

	// Run HTTP server in a goroutine
	go func() {
		log.Printf("INFO: HTTP server started on %s\n", s.cfg.HttpServerAddress)
		if err := server.ListenAndServeTLS(certFile, certKey); err != nil {
			log.Fatal("Could not start HTTP server:", err)
		}
	}()

	// WebSocket handling via Chi router
	s.router.HandleFunc("/ws", peer.p2p.wsHandler)

	log.Printf("INFO: WS server started on %s\n", s.cfg.WsServerAddress)
	if err := http.ListenAndServeTLS(s.cfg.WsServerAddress, certFile, certKey, nil); err != nil {
		log.Fatal("Could not start WebSocket server:", err)
	}
}

func (s *HttpApplication) AddPost(endpoint string, handler http.HandlerFunc) *HttpApplication {
	s.router.Post(endpoint, handler)
	return s

}

func (s *HttpApplication) AddGet(endpoint string, handler http.HandlerFunc) *HttpApplication {
	s.router.Get(endpoint, handler)
	return s
}
