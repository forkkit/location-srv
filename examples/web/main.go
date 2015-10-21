package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"github.com/myodc/go-micro/cmd"
	"github.com/myodc/go-micro/client"

	common "github.com/myodc/geo-srv/proto"
	search "github.com/myodc/geo-srv/proto/location/search"
	save "github.com/myodc/geo-srv/proto/location/save"
)

type Feature struct {
	Type string
	Properties map[string]interface{}
	Geometry Geometry
}

type Geometry struct {
	Type string
	Coordinates [][2]float64
}

type Route struct {
	Type string
	Features []Feature
}

func requestEntity(typ string, num int64, radius, lat, lon float64) ([]*common.Entity, error) {
	req := client.NewRequest("go.micro.srv.geo", "Location.Search", &search.Request{
		Center: &common.Location{
			Latitude: lat,
			Longitude: lon,
		},
		Type: typ,
		Radius: radius,
		NumEntities: num,
	})

	rsp := &search.Response{}
	err := client.Call(context.Background(), req, rsp)
	if err != nil {
		return nil, err
	}
	return rsp.Entities, nil
}

func saveEntity(id, typ string, lat, lon float64) {
        entity := &common.Entity{
                Id:   id,
                Type: typ,
                Location: &common.Location{
                        Latitude:  lat,
                        Longitude: lon,
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
}

func runner() {
	b, _ := ioutil.ReadFile("routes/strand.json")
	var route Route
	err := json.Unmarshal(b, &route)
	if err != nil {
		fmt.Errorf("erroring read route: %v", err)
		return
	}

	coords := route.Features[0].Geometry.Coordinates

	for {
		for i := 0; i < len(coords); i++ {
			saveEntity("one", "runner", coords[i][1], coords[i][0])
			time.Sleep(time.Second)
		}

		for i := len(coords) - 1; i >= 0; i-- {
			saveEntity("one", "runner", coords[i][1], coords[i][0])
			time.Sleep(time.Second)
		}
	}
}


func objectHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lat, _ := strconv.ParseFloat(r.Form.Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.Form.Get("lon"), 64)

	e, err := requestEntity("runner", 10, 500.0, lat, lon)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	b, err := json.Marshal(e)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(b)
}

func main() {
	cmd.Init()

	go runner()

	http.Handle("/", http.FileServer(http.Dir("html")))
	http.HandleFunc("/objects", objectHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
