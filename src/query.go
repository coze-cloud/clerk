package clerk

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type query[T any] struct {
	collection *Collection
	filter     map[string]any
	sorting    bson.D
	skip       int
	take       int
}

func NewQuery[T any](collection *Collection) *query[T] {
	return &query[T]{
		collection: collection,
		filter:     map[string]any{},
		sorting:    bson.D{},
		skip:       -1,
		take:       -1,
	}
}

func (q *query[T]) Where(key string, value any) *query[T] {
	q.filter[key] = value
	return q
}

func (q *query[T]) SortBy(key string, asc bool) *query[T] {
	var direction int
	if asc {
		direction = 1
	} else {
		direction = -1
	}

	q.sorting = append(q.sorting, bson.E{
		Key:   key,
		Value: direction,
	})
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
