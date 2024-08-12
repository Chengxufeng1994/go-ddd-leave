package aggregate

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/event"
)

type Aggregate interface {
	entity.Entity
	GetVersion() int64
	SetVersion(int64)
	GetEvents() []event.Event
	AddEvents(event event.Event) error
	ApplyEvent(event event.Event) error
	ClearEvents()
}

type AggregateBase struct {
	*entity.EntityBase
	id      string
	version int64
	events  []event.Event // 領域事件列表, 用於臨時存放完成某個業務流程中所發出的事件
}

func NewAggregateBase() *AggregateBase {
	entity := entity.NewEntityBase()
	return &AggregateBase{
		EntityBase: entity,
		events:     make([]event.Event, 0),
	}
}

func (agg *AggregateBase) GetID() string {
	return agg.id
}

func (agg *AggregateBase) SetID(id string) {
	agg.id = id
}

func (agg *AggregateBase) GetVersion() int64 {
	return agg.version
}

func (agg *AggregateBase) SetVersion(version int64) {
	agg.version = version
}

func (agg *AggregateBase) GetEvents() []event.Event {
	return agg.events
}

func (agg *AggregateBase) AddEvents(event event.Event) error {
	agg.events = append(agg.events, event)
	return nil
}

func (agg *AggregateBase) ApplyEvent(event event.Event) error {
	return nil
}

func (agg *AggregateBase) ClearEvents() {
	agg.events = make([]event.Event, 0)
}
