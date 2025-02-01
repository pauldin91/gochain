package main

import (
	"log"

	"github.com/pauldin91/gochain/src/api"
	"github.com/pauldin91/gochain/src/internal"
)

func main() {
	sb := api.ServerBuilder{Server: &api.Server{}}
	cfg, err := internal.LoadConfig(".")
	if err != nil {
		log.Fatal("unable to load config")
	}
	server := sb.
		WithConfig(cfg).
		WithRouter().
		Build()
	server.Start()
}
