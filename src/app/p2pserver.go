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
}

func (s *WsServer) connectToPeers(peers string) {

	for _, p := range strings.Split(peers, ",") {
		done := make(chan bool)
		go connect(p, done)
	}
}

func connect(peer string, done chan bool) {
	c, _, err := websocket.DefaultDialer.Dial(peer, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil && err != websocket.ErrCloseSent {
			log.Println("read:", err)
			done <- true
			break
		} else if err != websocket.ErrCloseSent {
			log.Println("Client")
			done <- true
			break
		}
		log.Printf("recv: %s", message)
	}

}
