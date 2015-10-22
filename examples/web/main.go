package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/myodc/go-micro/client"
	"github.com/myodc/go-micro/cmd"
	"golang.org/x/net/context"

	common "github.com/myodc/geo-srv/proto"
	save "github.com/myodc/geo-srv/proto/location/save"
	search "github.com/myodc/geo-srv/proto/location/search"
)

type Feature struct {
	Type       string
	Properties map[string]interface{}
	Geometry   Geometry
}

type Geometry struct {
	Type        string
	Coordinates [][2]float64
}

type Route struct {
	Type     string
	Features []Feature
}

func requestEntity(typ string, num int64, radius, lat, lon float64) ([]*common.Entity, error) {
	req := client.NewRequest("go.micro.srv.geo", "Location.Search", &search.Request{
		Center: &common.Location{
			Latitude:  lat,
			Longitude: lon,
		},
		Type:        typ,
		Radius:      radius,
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

func start() {
	b, _ := ioutil.ReadFile("routes/strand.json")
	var route Route
	err := json.Unmarshal(b, &route)
	if err != nil {
		fmt.Errorf("erroring read route: %v", err)
		return
	}

	coords := route.Features[0].Geometry.Coordinates
	go runner("one", "runner", coords)

	for i := 0; i < 20; i++ {
		time.Sleep(time.Second * 30)
		go runner(fmt.Sprintf("%d", time.Now().Unix()), "runner", coords)
	}
}

func runner(id, typ string, coords [][2]float64) {
	for {
		for i := 0; i < len(coords); i++ {
			saveEntity(id, typ, coords[i][1], coords[i][0])
			time.Sleep(time.Second)
		}

		for i := len(coords) - 1; i >= 0; i-- {
			saveEntity(id, typ, coords[i][1], coords[i][0])
			time.Sleep(time.Second)
		}
	}
}

func objectHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lat, _ := strconv.ParseFloat(r.Form.Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.Form.Get("lon"), 64)

	e, err := requestEntity("runner", 10, 1500.0, lat, lon)
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

	go start()

	http.Handle("/", http.FileServer(http.Dir("html")))
	http.HandleFunc("/objects", objectHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
