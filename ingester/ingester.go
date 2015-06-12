package ingester

import (
	log "github.com/golang/glog"
	"github.com/myodc/geo-srv/dao"
	"github.com/myodc/geo-srv/domain"
	proto "github.com/myodc/geo-srv/proto"
	"golang.org/x/net/context"
)

var (
	Topic = "geo.location"
)

type Geo struct{}

func (g *Geo) Handle(ctx context.Context, e *proto.Entity) error {
	log.Infof("Saving entity ID %s", e.Id)
	dao.Save(domain.ProtoToEntity(e))
	return nil
}
