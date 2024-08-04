package event

import (
	"time"

	"github.com/google/uuid"
)

type DomainEventName string
type DomainEventType string
type DomainEventStatus string

const (
	CREATED         DomainEventStatus = "CREATED"
	PUBLISH_SUCCEED DomainEventStatus = "PUBLISH_SUCCEED"
	PUBLISH_FAILED  DomainEventStatus = "PUBLISH_FAILED"
)

type BaseDomainEvent struct {
	id        string
	name      DomainEventName
	eType     DomainEventType
	status    DomainEventStatus
	timestamp time.Time
	data      []byte
}

// Type implements Event.

func NewBaseDomainEvent(eType DomainEventType, data []byte) BaseDomainEvent {
	return BaseDomainEvent{
		id:        uuid.New().String(),
		eType:     eType,
		status:    CREATED,
		timestamp: time.Now(),
		data:      data,
	}
}

func (event *BaseDomainEvent) ID() string {
	return event.id
}

func (event *BaseDomainEvent) Name() DomainEventName {
	return event.name
}

func (event *BaseDomainEvent) Type() DomainEventType {
	return event.eType
}

func (event *BaseDomainEvent) Status() DomainEventStatus {
	return event.status
}

func (event *BaseDomainEvent) Time() time.Time {
	return event.timestamp
}

func (event *BaseDomainEvent) Data() []byte {
	return event.data
}

type Event interface {
	ID() string
	Name() DomainEventName
	Type() DomainEventType
	Status() DomainEventStatus
	Time() time.Time
	Data() []byte
}

type EventHandler interface {
	Handle(Event) error
}

type EventPublisher interface {
	Subscribe(EventHandler, ...Event)
	Publish(Event)
}
