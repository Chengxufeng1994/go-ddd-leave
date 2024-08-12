package aggregate

import "context"

type AggregateRepository interface {
	Load(ctx context.Context, aggregateID string) Aggregate
	Save(ctx context.Context, aggregate Aggregate) (Aggregate, error)
	Exists(ctx context.Context, aggregateID string) (bool, error)
}
