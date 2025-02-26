package main

import "github.com/pauldin91/gochain/src/app"

func main() {
	s := app.NewServerBuilder()

	httpServer := s.
		WithConfig(".").
		WithRouter().
		WithWsServer().
		Build()

	httpServer.Start()
}
