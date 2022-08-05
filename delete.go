package clerk

import "context"

type delete[T any] struct {
	collection *Collection
	filter     map[string]any
}

func NewDelete[T any](collection *Collection) *delete[T] {
	return &delete[T]{
		collection: collection,
		filter:     map[string]any{},
	}
}

func (d *delete[T]) Where(key string, value any) *delete[T] {
	d.filter[key] = value
	return d
}

func (d *delete[T]) Execute(ctx context.Context, deleter Deleter[T]) error {
	return deleter.Delete(ctx, d.collection, d.filter)
}
