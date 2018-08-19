package dao

import (
	"sync"

	geo "github.com/hailocab/go-geoindex"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"
	"github.com/microhq/geo-srv/domain"
)

var (
	mtx          sync.RWMutex
	defaultIndex = geo.NewPointsIndex(geo.Km(0.5))
)

func Read(id string) (*domain.Entity, error) {
	mtx.RLock()
	defer mtx.RUnlock()

	p := defaultIndex.Get(id)
	if p == nil {
		return nil, errors.NotFound(server.DefaultOptions().Name+".read", "Not found")
	}

	entity, ok := p.(*domain.Entity)
	if !ok {
		return nil, errors.InternalServerError(server.DefaultOptions().Name+".read", "Error reading entity")
	}

	return entity, nil
}

func Save(e *domain.Entity) {
	mtx.Lock()
	defaultIndex.Add(e)
	mtx.Unlock()
}

func Search(typ string, entity *domain.Entity, radius float64, numEntities int) []*domain.Entity {
	mtx.RLock()
	defer mtx.RUnlock()

	points := defaultIndex.KNearest(entity, numEntities, geo.Meters(radius), func(p geo.Point) bool {
		e, ok := p.(*domain.Entity)
		if !ok || e.Type != typ {
			return false
		}
		return true
	})

	var entities []*domain.Entity

	for _, point := range points {
		e, ok := point.(*domain.Entity)
		if !ok {
			continue
		}
		entities = append(entities, e)
	}

	return entities
}
