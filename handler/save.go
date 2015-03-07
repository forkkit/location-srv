package handler

import (
	"code.google.com/p/go.net/context"

	save "github.com/asim/geo-srv/proto/location/save"
	"github.com/asim/go-micro/errors"
	"github.com/asim/go-micro/server"
	log "github.com/golang/glog"
)

func (l *Location) Save(ctx context.Context, req *save.Request, rsp *save.Response) error {
	log.Info("Received Location.Save request")

	entity := req.GetEntity()

	if entity.GetLocation() == nil {
		return errors.BadRequest(server.Name+".save", "Require location")
	}

	l.Index.Add(&Entity{
		id:        entity.GetId(),
		typ:       entity.GetType(),
		latitude:  entity.GetLocation().GetLatitude(),
		longitude: entity.GetLocation().GetLongitude(),
		timestamp: entity.GetLocation().GetTimestamp(),
	})

	return nil
}
