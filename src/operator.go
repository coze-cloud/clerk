package clerk

import "context"

type Creator[T any] interface {
	Create(
		ctx context.Context,
		collection *Collection,
		data T,
	) error
}

type Updater[T any] interface {
	Update(
		ctx context.Context,
		collection *Collection,
		filter map[string]any,
		data T,
		upsert bool,
	) error
}

type Deleter[T any] interface {
	Delete(
		ctx context.Context,
		collection *Collection,
		filter map[string]any,
	) error
}

type Querier[T any] interface {
	Query(
		ctx context.Context,
		collection *Collection,
		filter map[string]any,
	) (<-chan T, error)
}

type Watcher[T any] interface {
	Watch(
		ctx context.Context,
		collection *Collection,
		operation Operation,
	) (<-chan T, error)
}
