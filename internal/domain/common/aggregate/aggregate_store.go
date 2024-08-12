package aggregate

import "context"

type AggregateStore interface {
	Load(ctx context.Context, aggregate Aggregate) error
	Save(ctx context.Context, aggregate Aggregate) error
}
