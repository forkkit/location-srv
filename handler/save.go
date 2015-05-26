package handler

import (
	"encoding/json"

	"code.google.com/p/go.net/context"
	log "github.com/golang/glog"
	"github.com/myodc/geo-srv/domain"
	"github.com/myodc/geo-srv/ingester"
	save "github.com/myodc/geo-srv/proto/location/save"
	"github.com/myodc/go-micro/broker"
	"github.com/myodc/go-micro/errors"
	"github.com/myodc/go-micro/server"
)

func (l *Location) Save(ctx context.Context, req *save.Request, rsp *save.Response) error {
	log.Info("Received Location.Save request")

	entity := req.GetEntity()

	if entity.GetLocation() == nil {
		return errors.BadRequest(server.Config().Name()+".save", "Require location")
	}

	b, err := json.Marshal(domain.ProtoToEntity(entity))
	if err != nil {
		return errors.InternalServerError(server.Config().Name()+".save", err.Error())
	}

	broker.Publish(ctx, ingester.Topic, b)

	return nil
}
