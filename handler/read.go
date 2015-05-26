package handler

import (
	"code.google.com/p/go.net/context"

	log "github.com/golang/glog"
	"github.com/myodc/geo-srv/dao"
	read "github.com/myodc/geo-srv/proto/location/read"
	"github.com/myodc/go-micro/errors"
	"github.com/myodc/go-micro/server"
)

type Location struct{}

func (l *Location) Read(ctx context.Context, req *read.Request, rsp *read.Response) error {
	log.Info("Received Location.Read request")

	id := req.Id

	if len(id) == 0 {
		return errors.BadRequest(server.Config().Name()+".read", "Require Id")
	}

	entity, err := dao.Read(id)
	if err != nil {
		return err
	}

	rsp.Entity = entity.ToProto()

	return nil
}
