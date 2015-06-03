package main

import (
	log "github.com/golang/glog"
	"github.com/myodc/geo-srv/handler"
	"github.com/myodc/geo-srv/ingester"
	"github.com/myodc/go-micro/cmd"
	"github.com/myodc/go-micro/server"
)

func main() {
	// optionally setup command line usage
	cmd.Init()

	// Initialise Server
	server.Init(
		server.Name("go.micro.srv.geo"),
	)

	// Register Handlers
	server.Handle(
		server.NewHandler(new(handler.Location)),
	)

	// Start the ingester
	ingester.Run()

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
