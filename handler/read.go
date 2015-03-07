package handler

import (
	"code.google.com/p/go.net/context"

	"github.com/asim/geo-srv/dao"
	read "github.com/asim/geo-srv/proto/location/read"
	"github.com/asim/go-micro/errors"
	"github.com/asim/go-micro/server"
	log "github.com/golang/glog"
)

type Location struct{}

func (l *Location) Read(ctx context.Context, req *read.Request, rsp *read.Response) error {
	log.Info("Received Location.Read request")

	id := req.GetId()

	if len(id) == 0 {
		return errors.BadRequest(server.Name+".read", "Require Id")
	}

	entity, err := dao.Read(id)
	if err != nil {
		return err
	}

	rsp.Entity = entity.ToProto()

	return nil
}
