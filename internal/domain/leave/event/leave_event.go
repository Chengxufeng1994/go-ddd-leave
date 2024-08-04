package event

import (
	"encoding/json"

	common "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/event"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"
)

type LeaveEvent interface {
	common.Event
	LeaveID() string
}

type LeaveCreatedEvent struct {
	common.BaseDomainEvent
	leaveID string
}

func NewLeaveCreatedEvent(leave *entity.Leave) LeaveEvent {
	data, _ := json.Marshal(leave)
	return &LeaveCreatedEvent{
		BaseDomainEvent: common.NewBaseDomainEvent(Create, data),
		leaveID:         leave.GetID(),
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

func NewLeaveRejectedEvent(leave *entity.Leave) LeaveEvent {
	return &LeaveRejectedEvent{
		BaseDomainEvent: common.NewBaseDomainEvent(Create, nil),
		leaveID:         leave.GetID(),
	}
}

func (l *LeaveRejectedEvent) Name() common.DomainEventName {
	return "leave.rejected"
}

// LeaveID implements LeaveEvent.
func (l *LeaveRejectedEvent) LeaveID() string {
	return l.leaveID
}
