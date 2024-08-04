package aggregate

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/entity"
)

type Aggregate interface {
	entity.Entity
}

type AggregateBase struct {
	entity.EntityBase
	id string
}

func (agg *AggregateBase) GetID() string {
	return agg.id
}

func (agg *AggregateBase) SetID(id string) {
	agg.id = id
}
