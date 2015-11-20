package main

import (
	log "github.com/golang/glog"
	"github.com/micro/geo-srv/handler"
	"github.com/micro/geo-srv/ingester"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
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

	// Register Subscriber
	server.Subscribe(
		server.NewSubscriber(ingester.Topic, new(ingester.Geo)),
	)

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
