package broker

import "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/event"

type LeaveEventPublisher struct {
	handlers map[event.DomainEventName][]event.EventHandler
}

func NewLeavePublisher() event.EventPublisher {
	return &LeaveEventPublisher{
		handlers: make(map[event.DomainEventName][]event.EventHandler),
	}
}

// Subscribe implements event.EventPublisher.
func (e *LeaveEventPublisher) Subscribe(evtHandler event.EventHandler, evts ...event.Event) {
	for _, evt := range evts {
		handlers := e.handlers[evt.Name()]
		handlers = append(handlers, evtHandler)
		e.handlers[evt.Name()] = handlers
	}
}

// Publish implements event.EventPublisher.
func (e *LeaveEventPublisher) Publish(evt event.Event) {
	for _, h := range e.handlers[evt.Name()] {
		h.Handle(evt)
	}
}
