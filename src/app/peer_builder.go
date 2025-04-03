package app

import (
	"github.com/pauldin91/gochain/src/domain"
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

func (b *PeerBuilder) Build() *Peer {
	return b.peer
}
