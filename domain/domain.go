package domain

import (
	common "github.com/myodc/geo-srv/proto"
)

type Entity struct {
	ID        string
	Type      string
	Latitude  float64
	Longitude float64
	Timestamp int64
}

func (e *Entity) Id() string {
	return e.ID
}

func (e *Entity) Lat() float64 {
	return e.Latitude
}

func (e *Entity) Lon() float64 {
	return e.Longitude
}

func (e *Entity) ToProto() *common.Entity {
	return &common.Entity{
		Id:   e.ID,
		Type: e.Type,
		Location: &common.Location{
			Latitude:  e.Latitude,
			Longitude: e.Longitude,
			Timestamp: e.Timestamp,
		},
	}
}

func ProtoToEntity(e *common.Entity) *Entity {
	return &Entity{
		ID:        e.Id,
		Type:      e.Type,
		Latitude:  e.Location.Latitude,
		Longitude: e.Location.Longitude,
		Timestamp: e.Location.Timestamp,
	}
}
