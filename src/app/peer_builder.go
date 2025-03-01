package app

import (
	"log"

	"github.com/pauldin91/gochain/src/domain"
	"github.com/pauldin91/gochain/src/internal"
)

type PeerBuilder struct {
	peer *Peer
}

func (b *PeerBuilder) WithChain() *PeerBuilder {
	wallet := domain.NewWallet(0.0)
	pool := domain.TransactionPool{}
	chain := domain.Create()
	b.peer = &Peer{
		wallet: &wallet,
		pool:   &pool,
		chain:  chain,
	}
	return b
}

func (b *PeerBuilder) WithConfig(settings string) *PeerBuilder {
	cfg, err := internal.LoadConfig(settings)
	if err != nil {
		log.Fatal("unable to load config")
	}
	b.peer.cfg = cfg
	return b
}

func (b *PeerBuilder) WithPeerServer() *PeerBuilder {
	b.peer.p2p = &WsServer{}
	return b

}

func (b *PeerBuilder) Build() *Peer {
	return b.peer
}
