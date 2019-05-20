package handler

import (
	"log"

	"github.com/microhq/location-srv/dao"
	"github.com/microhq/location-srv/domain"
	loc "github.com/microhq/location-srv/proto/location"

	"golang.org/x/net/context"
)

func (l *Location) Search(ctx context.Context, req *loc.SearchRequest, rsp *loc.SearchResponse) error {
	log.Print("Received Location.Search request")

	entity := &domain.Entity{
		Latitude:  req.Center.Latitude,
		Longitude: req.Center.Longitude,
	}

	entities := dao.Search(req.Type, entity, req.Radius, int(req.NumEntities))

	for _, e := range entities {
		rsp.Entities = append(rsp.Entities, e.ToProto())
	}

	return nil
}
