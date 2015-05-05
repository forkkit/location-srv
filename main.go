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

	server.Name = "go.micro.srv.geo"

	// Initialise Server
	server.Init()

	// Register Handlers
	server.Register(
		server.NewReceiver(new(handler.Location)),
	)

	// Start the ingester
	ingester.Run()

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
