package main

import (
	"fmt"
	"time"

	"code.google.com/p/goprotobuf/proto"
	common "github.com/asim/geo-srv/proto"
	read "github.com/asim/geo-srv/proto/location/read"
	save "github.com/asim/geo-srv/proto/location/save"
	search "github.com/asim/geo-srv/proto/location/search"
	"github.com/asim/go-micro/client"
)

func saveEntity() {
	entity := &common.Entity{
		Id:   proto.String("id123"),
		Type: proto.String("runner"),
		Location: &common.Location{
			Latitude:  proto.Float64(51.516509),
			Longitude: proto.Float64(0.124615),
			Timestamp: proto.Int64(time.Now().Unix()),
		},
	}

	req := client.NewRequest("go.micro.srv.geo", "Location.Save", &save.Request{
		Entity: entity,
	})

	rsp := &save.Response{}

	if err := client.Call(req, rsp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Saved entity: %+v\n", entity)
}

func readEntity() {
	req := client.NewRequest("go.micro.srv.geo", "Location.Read", &read.Request{
		Id: proto.String("id123"),
	})

	rsp := &read.Response{}

	if err := client.Call(req, rsp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Read entity: %+v\n", rsp.GetEntity())
}

func searchForEntities() {
	req := client.NewRequest("go.micro.srv.geo", "Location.Search", &search.Request{
		Center: &common.Location{
			Latitude:  proto.Float64(51.516509),
			Longitude: proto.Float64(0.124615),
			Timestamp: proto.Int64(time.Now().Unix()),
		},
		Radius:      proto.Float64(500.0),
		Type:        proto.String("runner"),
		NumEntities: proto.Int64(5),
	})

	rsp := &search.Response{}

	if err := client.Call(req, rsp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Search results: %+v\n", rsp.GetEntities())

}

func main() {
	saveEntity()
	readEntity()
	searchForEntities()
}
