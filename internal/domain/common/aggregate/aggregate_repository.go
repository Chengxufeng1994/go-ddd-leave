package aggregate

import "context"

type AggregateRepository interface {
	Load(ctx context.Context, aggregateID string) Aggregate
}
