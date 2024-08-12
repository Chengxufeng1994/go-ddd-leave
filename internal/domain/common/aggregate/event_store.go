package aggregate

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/event"
)

type EventStore interface {
	SaveEvents(ctx context.Context, streamID string, events []event.Event) error
	LoadEvents(ctx context.Context, streamID string) ([]event.Event, error)
}
