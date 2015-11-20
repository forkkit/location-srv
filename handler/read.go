package handler

import (
	log "github.com/golang/glog"
	"github.com/micro/geo-srv/dao"
	read "github.com/micro/geo-srv/proto/location/read"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"

	"golang.org/x/net/context"
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
