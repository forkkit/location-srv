package handler

import (
	"encoding/json"

	"code.google.com/p/go.net/context"
	"github.com/asim/geo-srv/domain"
	"github.com/asim/geo-srv/ingester"
	save "github.com/asim/geo-srv/proto/location/save"
	"github.com/asim/go-micro/broker"
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

	b, err := json.Marshal(domain.ProtoToEntity(entity))
	if err != nil {
		return errors.InternalServerError(server.Name+".save", err.Error())
	}

	broker.Publish(ingester.Topic, b)

	return nil
}
