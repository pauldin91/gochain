package main

import "github.com/pauldin91/gochain/src/app"

func main() {
	peer := (&app.PeerBuilder{}).
		WithChain().
		WithConfig(".").
		WithPeerServer().
		Build()

	s := app.NewServerBuilder(peer)

	httpServer := s.
		WithConfig(".").
		WithRouter().
		Build()

	httpServer.Start(peer)
}
