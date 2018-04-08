package main

import (
	"fmt"
	"time"

	common "github.com/micro/geo-srv/proto"
	loc "github.com/micro/geo-srv/proto/location"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"

	"golang.org/x/net/context"
)

var (
	cl loc.LocationService
)

func saveEntity() {
	entity := &common.Entity{
		Id:   "id123",
		Type: "runner",
		Location: &common.Point{
			Latitude:  51.516509,
			Longitude: 0.124615,
			Timestamp: time.Now().Unix(),
		},
	}

	_, err := cl.Save(context.Background(), &loc.SaveRequest{
		Entity: entity,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Saved entity: %+v\n", entity)
}

func readEntity() {
	rsp, err := cl.Read(context.Background(), &loc.ReadRequest{
		Id: "id123",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Read entity: %+v\n", rsp.Entity)
}

func searchForEntities() {
	rsp, err := cl.Search(context.Background(), &loc.SearchRequest{
		Center: &common.Point{
			Latitude:  51.516509,
			Longitude: 0.124615,
			Timestamp: time.Now().Unix(),
		},
		Radius:      500.0,
		Type:        "runner",
		NumEntities: 5,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Search results: %+v\n", rsp.Entities)

}

func main() {
	// init flags
	cmd.Init()

	// use client stub
	cl = loc.LocationServiceClient("go.micro.srv.geo", client.DefaultClient)

	// do requests
	saveEntity()
	readEntity()
	searchForEntities()
}
