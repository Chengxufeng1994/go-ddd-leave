package event

import (
	"encoding/json"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/aggregate"
	common "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/event"
)

const (
	LeaveCreated = "LEAVE_CREATED"
)

type LeaveEvent interface {
	common.Event
	LeaveID() string
}

type LeaveCreatedEvent struct {
	common.BaseDomainEvent
	leaveID string
}

func NewLeaveCreatedEvent(aggregate aggregate.Aggregate) LeaveEvent {
	data, _ := json.Marshal(aggregate)
	return &LeaveCreatedEvent{
		BaseDomainEvent: common.NewBaseDomainEvent(Create, data),
		leaveID:         aggregate.GetID(),
	}
}

func (event *LeaveCreatedEvent) Name() common.DomainEventName {
	return "leave.created"
}

func (event *LeaveCreatedEvent) LeaveID() string {
	return event.leaveID
}

type LeaveRejectedEvent struct {
	common.BaseDomainEvent
	leaveID string
}

func NewLeaveRejectedEvent(aggregate aggregate.Aggregate) LeaveEvent {
	return &LeaveRejectedEvent{
		BaseDomainEvent: common.NewBaseDomainEvent(Reject, nil),
		leaveID:         aggregate.GetID(),
	}
}

func (l *LeaveRejectedEvent) Name() common.DomainEventName {
	return "leave.rejected"
}

// LeaveID implements LeaveEvent.
func (l *LeaveRejectedEvent) LeaveID() string {
	return l.leaveID
}
