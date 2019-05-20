package main

import (
	"log"
	"time"

	"github.com/micro/go-micro"
	"github.com/microhq/location-srv/handler"
	"github.com/microhq/location-srv/ingester"
	proto "github.com/microhq/location-srv/proto/location"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.location"),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*30),
	)

	// Initialise Server
	service.Init()

	// Register Handlers
	proto.RegisterLocationHandler(service.Server(), new(handler.Location))

	// Register Subscriber
	micro.RegisterSubscriber(ingester.Topic, service.Server(), new(ingester.Geo))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
