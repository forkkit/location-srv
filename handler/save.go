package handler

import (
	"log"

	"github.com/micro/geo-srv/ingester"
	loc "github.com/micro/geo-srv/proto/location"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"

	"golang.org/x/net/context"
)

func (l *Location) Save(ctx context.Context, req *loc.SaveRequest, rsp *loc.SaveResponse) error {
	log.Print("Received Location.Save request")

	entity := req.GetEntity()

	if entity.GetLocation() == nil {
		return errors.BadRequest(server.DefaultOptions().Name+".save", "Require location")
	}

	p := loc.NewPublisher(ingester.Topic, client.DefaultClient)
	if err := p.Publish(ctx, entity); err != nil {
		return errors.InternalServerError(server.DefaultOptions().Name+".save", err.Error())
	}

	return nil
}
