package app

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pauldin91/gochain/src/domain"
	"github.com/pauldin91/gochain/src/utils"
)

var chain domain.Blockchain

type WsServer struct {
	sockets map[string]*websocket.Conn
	mutex   sync.Mutex
	cfg     utils.Config
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type BroadcastMessage struct {
	Message string `json:"message"`
}

func (ws *WsServer) wsHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade connection", http.StatusBadRequest)
		return
	}
	clientID := fmt.Sprintf("%p", conn)

	ws.mutex.Lock()
	ws.sockets[clientID] = conn
	ws.mutex.Unlock()

	fmt.Println("New client connected with id : ", clientID)

	var msg BroadcastMessage = BroadcastMessage{

		Message: fmt.Sprintf("You are connected with id %s", clientID),
	}
	ws.sendMessageToClient(clientID, msg)
	ws.sendMessageToClient(clientID, chain)

	defer ws.closeConnection(clientID)

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		ws.broadcastMessage(fmt.Sprintf("clientid %s says %s\n", clientID, string(p)))
	}
}

func (ws *WsServer) closeConnection(clientID string) {
	ws.mutex.Lock()
	ws.sockets[clientID].Close()
	delete(ws.sockets, clientID)
	ws.mutex.Unlock()
	fmt.Println("Client disconnected:", clientID)
}

func (ws *WsServer) broadcastMessage(message string) {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	for clientID, conn := range ws.sockets {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println("Error writing to client", clientID, ":", err)
			conn.Close()
			delete(ws.sockets, clientID)
		}
	}
}

func (ws *WsServer) sendMessageToClient(clientID string, message any) {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	if conn, ok := ws.sockets[clientID]; ok {
		if err := conn.WriteJSON(message); err != nil {
			log.Println("Error sending to client", clientID, ":", err)
			conn.Close()
			delete(ws.sockets, clientID)
		}
	}
}

func (s *WsServer) connectToPeers(peer string) {
	var peers []string = strings.Split(peer, ",")
	var chans []chan bool = make([]chan bool, len(peers))
	for i, p := range peers {
		chans[i] = make(chan bool)
		go s.connect(p, chans[i])
	}

	for _, c := range chans {
		<-c
	}
}

func (s *WsServer) connect(peer string, done chan bool) {
	if peer == "" {
		return
	}
	dialer := websocket.Dialer{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	c, _, err := dialer.Dial(peer, nil)
	if err != nil {
		log.Printf("dial %s error \n", peer)
		done <- true
		return
	}

	log.Printf("connected to server %s", peer)

	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			done <- true
			break
		}
		log.Printf("recv: %s", message)
	}
}
