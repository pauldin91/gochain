package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pauldin91/gochain/src/domain"
	"github.com/pauldin91/gochain/src/internal"
)

type Peer struct {
	p2p    *WsServer
	wallet *domain.Wallet
	pool   *domain.TransactionPool
	chain  *domain.Blockchain
	cfg    internal.Config
}

func (ner *Peer) mine() domain.Block {
	validTransactions := ner.pool.ValidTransactions()
	validTransactions = append(validTransactions, *domain.Reward(ner.wallet, ner.wallet))

	data, _ := json.Marshal(validTransactions)
	block := ner.chain.AddBlock(string(data))

	ner.syncChains()
	ner.pool.Clear()
	ner.broadcast(ner.chain.Chain)

	return block
}

func (ner *Peer) broadcast(chain []domain.Block) {
	for _, ws := range ner.p2p.sockets {
		err := ws.WriteJSON(chain)
		if err != nil {
			log.Println("error writing", err)
		}
	}
}

func (ner *Peer) Clear() {
	ner.pool.Clear()
}

func (ner *Peer) syncChains() {

	chain, _ := json.Marshal(ner.chain)
	for _, s := range ner.p2p.sockets {
		s.WriteJSON(string(chain))
	}
}

func (peer *Peer) listen(w http.ResponseWriter, req *http.Request) {
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

	if peer.p2p == nil {
		peer.p2p = &WsServer{}
	}
	clientId := fmt.Sprintf("%p", ws)
	peer.p2p.sockets[clientId] = ws
	peer.broadcast(chain.Chain)
}
