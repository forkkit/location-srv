package main

import (
	"fmt"
	"time"

	common "github.com/myodc/geo-srv/proto"
	read "github.com/myodc/geo-srv/proto/location/read"
	save "github.com/myodc/geo-srv/proto/location/save"
	search "github.com/myodc/geo-srv/proto/location/search"
	"github.com/myodc/go-micro/client"
	"github.com/myodc/go-micro/cmd"

	"golang.org/x/net/context"
)

func saveEntity() {
	entity := &common.Entity{
		Id:   "id123",
		Type: "runner",
		Location: &common.Location{
			Latitude:  51.516509,
			Longitude: 0.124615,
			Timestamp: time.Now().Unix(),
		},
	}

	req := client.NewRequest("go.micro.srv.geo", "Location.Save", &save.Request{
		Entity: entity,
	})

	rsp := &save.Response{}

	if err := client.Call(context.Background(), req, rsp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Saved entity: %+v\n", entity)
}

func readEntity() {
	req := client.NewRequest("go.micro.srv.geo", "Location.Read", &read.Request{
		Id: "id123",
	})

	rsp := &read.Response{}

	if err := client.Call(context.Background(), req, rsp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Read entity: %+v\n", rsp.Entity)
}

func searchForEntities() {
	req := client.NewRequest("go.micro.srv.geo", "Location.Search", &search.Request{
		Center: &common.Location{
			Latitude:  51.516509,
			Longitude: 0.124615,
			Timestamp: time.Now().Unix(),
		},
		Radius:      500.0,
		Type:        "runner",
		NumEntities: 5,
	})

	rsp := &search.Response{}

	if err := client.Call(context.Background(), req, rsp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Search results: %+v\n", rsp.Entities)

}

func main() {
	cmd.Init()
	saveEntity()
	readEntity()
	searchForEntities()
}
