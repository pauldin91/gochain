package main

import "github.com/pauldin91/gochain/src/app"

func main() {
	//peer := (&app.PeerBuilder{}).
	//	WithChain().
	//	Build()
	//
	s := app.NewServerBuilder()

	httpServer := s.
		WithConfig(".").
		WithRouter().
		WithPeerServer().
		Build()

	httpServer.Start()
}
