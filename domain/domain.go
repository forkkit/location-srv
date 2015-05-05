package domain

import (
	"github.com/golang/protobuf/proto"
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
		Id:   proto.String(e.ID),
		Type: proto.String(e.Type),
		Location: &common.Location{
			Latitude:  proto.Float64(e.Latitude),
			Longitude: proto.Float64(e.Longitude),
			Timestamp: proto.Int64(e.Timestamp),
		},
	}
}

func ProtoToEntity(e *common.Entity) *Entity {
	return &Entity{
		ID:        e.GetId(),
		Type:      e.GetType(),
		Latitude:  e.GetLocation().GetLatitude(),
		Longitude: e.GetLocation().GetLongitude(),
		Timestamp: e.GetLocation().GetTimestamp(),
	}
}
