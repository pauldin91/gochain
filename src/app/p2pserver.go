package app

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/pauldin91/gochain/src/internal"
)

type WsServer struct {
	cfg     internal.Config
	sockets []*websocket.Conn
}

type WsServerBuilder struct {
	Server *WsServer
}

func (sb WsServerBuilder) Build() *WsServer {
	return sb.Server
}

func (sb WsServerBuilder) WithConfig(settings string) WsServerBuilder {
	cfg, err := internal.LoadConfig(settings)
	if err != nil {
		log.Fatal("unable to load config")
	}
	sb.Server.cfg = cfg
	return sb
}

type SimpleRequest struct {
	Message string `json:"msg"`
}

type SimpleResponse struct {
	Response string `json:"response"`
}

var received []SimpleRequest

func (p *WsServer) Start() {

	certFile := filepath.Join(p.cfg.CertPath, p.cfg.CertFile)
	certKey := filepath.Join(p.cfg.CertPath, p.cfg.CertKey)
	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		log.Fatal("unable to load certs")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		upgrader := websocket.Upgrader{}
		ws, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Printf("could not establish a connection with %s", ws.RemoteAddr())
		}

		ws.SetReadLimit(p.cfg.WsReadLimit)
		p.sockets = append(p.sockets, ws)
		for {
			simple := SimpleRequest{}
			err := ws.ReadJSON(&simple)
			if err != nil {
				log.Println("read:", err)
				break
			}
			received = append(received, simple)
			p.broadcast()
		}
	})

	err := http.ListenAndServeTLS(p.cfg.WsServerAddress, certFile, certKey, nil)
	if err != nil {
		log.Fatal("Could not start ws server", err)
	}

}

func (s *WsServer) broadcast() {
	for _, ws := range s.sockets {
		err := ws.WriteJSON(received)
		if err != nil {
			log.Println("error writing", err)
		}
	}
}

func (s *WsServer) connectToPeers() {

	for _, p := range strings.Split(s.cfg.Peers, ",") {
		go connect(p)
	}
}

func connect(peer string) {
	c, _, err := websocket.DefaultDialer.Dial(peer, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}
