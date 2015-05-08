package handler

import (
	"code.google.com/p/go.net/context"

	log "github.com/golang/glog"
	"github.com/myodc/geo-srv/dao"
	"github.com/myodc/geo-srv/domain"
	search "github.com/myodc/geo-srv/proto/location/search"
)

func (l *Location) Search(ctx context.Context, req *search.Request, rsp *search.Response) error {
	log.Info("Received Location.Search request")

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
