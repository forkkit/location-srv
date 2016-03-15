package main

import (
	"log"
	"time"

	"github.com/micro/geo-srv/handler"
	"github.com/micro/geo-srv/ingester"
	proto "github.com/micro/geo-srv/proto/location"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.geo"),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*30),
	)

	// Initialise Server
	service.Init()

	// Register Handlers
	proto.RegisterLocationHandler(service.Server(), new(handler.Location))

	// Register Subscriber
	service.Server().Subscribe(
		service.Server().NewSubscriber(
			ingester.Topic,
			new(ingester.Geo),
		),
	)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
