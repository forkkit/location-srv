package main

import (
	"github.com/asim/geo-srv/handler"
	"github.com/asim/go-micro/cmd"
	"github.com/asim/go-micro/server"
	log "github.com/golang/glog"
	geo "github.com/hailocab/go-geoindex"
)

func main() {
	// optionally setup command line usage
	cmd.Init()

	server.Name = "go.micro.srv.geo"

	// Initialise Server
	server.Init()

	// Register Handlers
	server.Register(
		server.NewReceiver(&handler.Location{
			Index: geo.NewPointsIndex(geo.Km(0.5)),
		}),
	)

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
