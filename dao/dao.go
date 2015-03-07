package dao

import (
	"github.com/asim/geo-srv/domain"
	"github.com/asim/go-micro/errors"
	"github.com/asim/go-micro/server"
	geo "github.com/hailocab/go-geoindex"
)

var (
	defaultIndex = geo.NewPointsIndex(geo.Km(0.5))
)

func Read(id string) (*domain.Entity, error) {
	p := defaultIndex.Get(id)
	if p == nil {
		return nil, errors.NotFound(server.Name+".read", "Not found")
	}

	entity, ok := p.(*domain.Entity)
	if !ok {
		return nil, errors.InternalServerError(server.Name+".read", "Error reading entity")
	}

	return entity, nil
}

func Save(e *domain.Entity) {
	defaultIndex.Add(e)
}

func Search(typ string, entity *domain.Entity, radius float64, numEntities int) []*domain.Entity {
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
