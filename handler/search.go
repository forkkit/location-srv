package handler

import (
	"code.google.com/p/go.net/context"

	"github.com/asim/geo-srv/dao"
	"github.com/asim/geo-srv/domain"
	search "github.com/asim/geo-srv/proto/location/search"
	log "github.com/golang/glog"
)

func (l *Location) Search(ctx context.Context, req *search.Request, rsp *search.Response) error {
	log.Info("Received Location.Search request")

	entity := &domain.Entity{
		Latitude:  req.GetCenter().GetLatitude(),
		Longitude: req.GetCenter().GetLongitude(),
	}

	entities := dao.Search(req.GetType(), entity, req.GetRadius(), int(req.GetNumEntities()))

	for _, e := range entities {
		rsp.Entities = append(rsp.Entities, e.ToProto())
	}

	return nil
}
