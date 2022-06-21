package clerk

import "context"

type query[T any] struct {
	collection *Collection
	filter     map[string]any
}

func NewQuery[T any](collection *Collection) *query[T] {
	return &query[T]{
		collection: collection,
		filter:     map[string]any{},
	}
}

func (q *query[T]) Where(key string, value any) *query[T] {
	q.filter[key] = value
	return q
}

func (q *query[T]) Execute(ctx context.Context, querier Querier[T]) (<-chan T, error) {
	return querier.Query(
		ctx,
		q.collection,
		q.filter,
	)
}
