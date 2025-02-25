package main

import "github.com/pauldin91/gochain/src/app"

func main() {
	s := app.ServerBuilder{&app.HttpServer{}}

	httpServer := s.
		WithConfig(".").
		WithRouter().
		WithWsServer().
		Build()

	httpServer.Start()
}
