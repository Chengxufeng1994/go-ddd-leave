package event

import "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/event"

const (
	Create   event.DomainEventType = "create"
	Agree    event.DomainEventType = "agree"
	Reject   event.DomainEventType = "reject"
	Approved event.DomainEventType = "approved"
)
