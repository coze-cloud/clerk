package clerk

import "context"

type update[T any] struct {
	collection *Collection
	filter     map[string]any
	data       T
	upsert     bool
}

func NewUpdate[T any](collection *Collection, data T) *update[T] {
	return &update[T]{
		collection: collection,
		filter:     map[string]any{},
		data:       data,
		upsert:     false,
	}
}

func (c *update[T]) Where(key string, value any) *update[T] {
	c.filter[key] = value
	return c
}

func (c *update[T]) WithUpsert() *update[T] {
	c.upsert = true
	return c
}

func (c *update[T]) Execute(ctx context.Context, updater Updater[T]) error {
	return updater.Update(ctx, c.collection, c.filter, c.data, c.upsert)
}
