package ingester

import (
	"encoding/json"
	"time"

	"github.com/asim/geo-srv/dao"
	"github.com/asim/geo-srv/domain"
	"github.com/asim/micro-mq/go/client"
	log "github.com/golang/glog"
)

var (
	Topic = "geo.location"
)

func Run() {
	for {
		ch, err := client.Subscribe(Topic)
		if err != nil {
			log.Error(err)
			continue
		}

		for e := range ch {
			var entity *domain.Entity
			if err := json.Unmarshal(e, &entity); err != nil {
				log.Warning(err.Error())
				continue
			}
			dao.Save(entity)
		}

		log.Warning("Subscription channel closed")
		time.Sleep(time.Second * 5)
	}
}
