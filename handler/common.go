package handler

import (
	geo "github.com/hailocab/go-geoindex"
)

type Entity struct {
	id        string
	typ       string
	latitude  float64
	longitude float64
	timestamp int64
}

type Location struct {
	Index *geo.PointsIndex
}

func (e *Entity) Id() string {
	return e.id
}

func (e *Entity) Lat() float64 {
	return e.latitude
}

func (e *Entity) Lon() float64 {
	return e.longitude
}
