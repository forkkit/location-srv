package handler

import (
	"code.google.com/p/go.net/context"
	"code.google.com/p/goprotobuf/proto"

	common "github.com/asim/geo-srv/proto"
	search "github.com/asim/geo-srv/proto/location/search"
	log "github.com/golang/glog"
	geo "github.com/hailocab/go-geoindex"
)

func (l *Location) Search(ctx context.Context, req *search.Request, rsp *search.Response) error {
	log.Info("Received Location.Search request")

	entity := &Entity{
		latitude:  req.GetCenter().GetLatitude(),
		longitude: req.GetCenter().GetLongitude(),
	}

	points := l.Index.KNearest(entity, int(req.GetNumEntities()), geo.Meters(req.GetRadius()), func(p geo.Point) bool {
		e, ok := p.(*Entity)
		if !ok || e.typ != req.GetType() {
			return false
		}
		return true
	})

	for _, point := range points {
		e, ok := point.(*Entity)
		if !ok {
			continue
		}

		rsp.Entities = append(rsp.Entities, &common.Entity{
			Id:   proto.String(e.id),
			Type: proto.String(e.typ),
			Location: &common.Location{
				Latitude:  proto.Float64(e.latitude),
				Longitude: proto.Float64(e.longitude),
				Timestamp: proto.Int64(e.timestamp),
			},
		})
	}

	return nil
}
