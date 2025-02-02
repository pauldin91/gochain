package app

import (
	"log"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/pauldin91/gochain/src/domain"
)

var chain domain.Blockchain

type WsServer struct {
	sockets []*websocket.Conn
	peers   string
}

func (s *WsServer) broadcast(chain []domain.Block) {
	for _, ws := range s.sockets {
		err := ws.WriteJSON(chain)
		if err != nil {
			log.Println("error writing", err)
		}
	}
}

func (s *WsServer) connectToPeers() {

	for _, p := range strings.Split(s.peers, ",") {
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
