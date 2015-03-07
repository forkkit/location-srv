package handler

import (
	"code.google.com/p/go.net/context"
	"code.google.com/p/goprotobuf/proto"

	common "github.com/asim/geo-srv/proto"
	read "github.com/asim/geo-srv/proto/location/read"
	"github.com/asim/go-micro/errors"
	"github.com/asim/go-micro/server"
	log "github.com/golang/glog"
)

func (l *Location) Read(ctx context.Context, req *read.Request, rsp *read.Response) error {
	log.Info("Received Location.Read request")

	id := req.GetId()

	if len(id) == 0 {
		return errors.BadRequest(server.Name+".read", "Require Id")
	}

	p := l.Index.Get(id)

	if p == nil {
		return errors.NotFound(server.Name+".read", "Not found")
	}

	entity, ok := p.(*Entity)
	if !ok {
		return errors.InternalServerError(server.Name+".read", "Error reading entity")
	}

	rsp.Entity = &common.Entity{
		Id:   proto.String(entity.id),
		Type: proto.String(entity.typ),
		Location: &common.Location{
			Latitude:  proto.Float64(entity.latitude),
			Longitude: proto.Float64(entity.longitude),
			Timestamp: proto.Int64(entity.timestamp),
		},
	}

	return nil
}
