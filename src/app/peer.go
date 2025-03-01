package app

import (
	"encoding/json"
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
	return block
}

func (ner *Peer) Clear() {
	ner.pool.Clear()
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

	peer.p2p.sockets = append(peer.p2p.sockets, ws)
	peer.p2p.broadcast(chain.Chain)
}
