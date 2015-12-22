package handler

import (
	log "github.com/golang/glog"
	"github.com/micro/geo-srv/ingester"
	loc "github.com/micro/geo-srv/proto/location"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"

	"golang.org/x/net/context"
)

func (l *Location) Save(ctx context.Context, req *loc.SaveRequest, rsp *loc.SaveResponse) error {
	log.Info("Received Location.Save request")

	entity := req.GetEntity()

	if entity.GetLocation() == nil {
		return errors.BadRequest(server.Config().Name()+".save", "Require location")
	}

	p := client.NewPublication(ingester.Topic, entity)

	if err := client.Publish(ctx, p); err != nil {
		return errors.InternalServerError(server.Config().Name()+".save", err.Error())
	}

	return nil
}
