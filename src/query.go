package clerk

import (
	"context"
)

type query[T any] struct {
	collection *Collection
	filter     map[string]any
	sorting    map[string]bool
	skip       int
	take       int
}

func NewQuery[T any](collection *Collection) *query[T] {
	return &query[T]{
		collection: collection,
		filter:     map[string]any{},
		sorting:    map[string]bool{},
		skip:       -1,
		take:       -1,
	}
}

func (q *query[T]) Where(key string, value any) *query[T] {
	q.filter[key] = value
	return q
}

func (q *query[T]) SortBy(key string, asc bool) *query[T] {
	q.sorting[key] = asc
	return q
}

func (q *query[T]) Skip(skip int) *query[T] {
	q.skip = skip
	return q
}

func (q *query[T]) Take(take int) *query[T] {
	q.take = take
	return q
}

func (q *query[T]) Execute(ctx context.Context, querier Querier[T]) (<-chan T, error) {
	return querier.Query(
		ctx,
		q.collection,
		q.filter,
		q.sorting,
		q.skip,
		q.take,
	)
}
