package ingester

import (
	"log"

	"github.com/microhq/location-srv/dao"
	"github.com/microhq/location-srv/domain"
	proto "github.com/microhq/location-srv/proto"
	"golang.org/x/net/context"
)

var (
	Topic = "geo.location"
)

type Geo struct{}

func (g *Geo) Handle(ctx context.Context, e *proto.Entity) error {
	log.Printf("Saving entity ID %s", e.Id)
	dao.Save(domain.ProtoToEntity(e))
	return nil
}
