package main

import (
	"github.com/pauldin91/gochain/src/app"
)

func main() {
	sb := app.WsServerBuilder{&app.WsServer{}}

	server := sb.
		WithConfig(".").
		Build()
	server.Start()
}
