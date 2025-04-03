package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pauldin91/gochain/src/domain"
)

type Peer struct {
	wallet *domain.Wallet
	pool   *domain.TransactionPool
	chain  *domain.Blockchain
}

func (ner *Peer) mine() domain.Block {
	validTransactions := ner.pool.ValidTransactions()
	validTransactions = append(validTransactions, *domain.Reward(ner.wallet, ner.wallet))

	data, _ := json.Marshal(validTransactions)
	block := ner.chain.AddBlock(string(data))

	//ner.syncChains()
	ner.pool.Clear()
	//ner.broadcast(ner.chain.Chain)

	return block
}

func (ner *Peer) Clear() {
	ner.pool.Clear()
}

func (ner *HttpApplication) syncChains() {

	chain, _ := json.Marshal(ner.peer.chain)
	for _, s := range ner.ws.sockets {
		s.WriteJSON(string(chain))
	}
}

func (peer *HttpApplication) listen(w http.ResponseWriter, req *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  int(peer.cfg.WsReadLimit),
		WriteBufferSize: int(peer.cfg.WsWriteLimit),
	}
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	ws.SetReadLimit(peer.cfg.WsReadLimit)

	if peer.ws == nil {
		peer.ws = &WsServer{}
	}
	clientId := fmt.Sprintf("%p", ws)
	peer.ws.sockets[clientId] = ws

	//peer.ws.broadcastMessage(chain.Chain.String())
}
