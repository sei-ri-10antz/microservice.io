package main

import (
	"context"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/sei-ri/microservice.io/account/server"
)

func main() {
	var srv server.Server
	if err := envconfig.Process("", &srv); err != nil {
		log.Fatal(err)
	}
	srv.Serve(context.Background())
}
