package app

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/pauldin91/gochain/src/domain"
	"github.com/pauldin91/gochain/src/internal"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: errMsg,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

type Server interface {
	Start()
}

type HttpServer struct {
	cfg    internal.Config
	chain  *domain.Blockchain
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

	// Run HTTP server in a goroutine
	go func() {
		log.Printf("INFO: HTTP server started on %s\n", s.cfg.HttpServerAddress)
		if err := server.ListenAndServeTLS(certFile, certKey); err != nil {
			log.Fatal("Could not start HTTP server:", err)
		}
	}()

	// WebSocket handling via Chi router
	s.router.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  int(s.cfg.WsReadLimit),
			WriteBufferSize: int(s.cfg.WsWriteLimit),
		}
		ws, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Printf("WebSocket upgrade failed: %v", err)
			return
		}

		ws.SetReadLimit(s.cfg.WsReadLimit)

		if s.p2p == nil {
			s.p2p = &WsServer{}
		}

		s.p2p.sockets = append(s.p2p.sockets, ws)
		s.p2p.broadcast(chain.Chain)
	})

	log.Printf("INFO: WS server started on %s\n", s.cfg.WsServerAddress)
	if err := http.ListenAndServeTLS(s.cfg.WsServerAddress, certFile, certKey, nil); err != nil {
		log.Fatal("Could not start WebSocket server:", err)
	}
}
