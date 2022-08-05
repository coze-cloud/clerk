package clerk

import "context"

type create[T any] struct {
	collection *Collection
	data       T
}

func NewCreate[T any](collection *Collection, data T) *create[T] {
	return &create[T]{
		collection: collection,
		data:       data,
	}
}

func (c *create[T]) Execute(ctx context.Context, creator Creator[T]) error {
	return creator.Create(ctx, c.collection, c.data)
}
