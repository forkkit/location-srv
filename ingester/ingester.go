package ingester

import (
	"encoding/json"

	log "github.com/golang/glog"
	"github.com/myodc/geo-srv/dao"
	"github.com/myodc/geo-srv/domain"
	"github.com/myodc/go-micro/broker"
	"golang.org/x/net/context"
)

var (
	Topic = "geo.location"
)

func Run() {
	log.Infof("Starting topic %s subscriber", Topic)
	broker.Init()
	if err := broker.Connect(); err != nil {
		log.Fatalf("Error connecting to broker: %v", err)
	}
	_, err := broker.Subscribe(Topic, func(ctx context.Context, msg *broker.Message) {
		var entity *domain.Entity
		if er := json.Unmarshal(msg.Body, &entity); er != nil {
			log.Warning(er.Error())
			return
		}
		log.Infof("Saving entity ID %s", entity.Id())
		dao.Save(entity)
	})
	if err != nil {
		log.Errorf("Error subscribing to topic %s: %v", Topic, err)
	}
}
