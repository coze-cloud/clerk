package clerk

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

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
		sorting bson.D,
		skip int,
		take int,
	) (<-chan T, error)
}

type Searcher[T any] interface {
	Search(
		ctx context.Context,
		collection *Collection,
		query string,
		highlight []string,
		filterable []string,
		filterQuery string,
		skip int,
		take int,
	) (<-chan T, error)
}

type Watcher[T any] interface {
	Watch(
		ctx context.Context,
		collection *Collection,
		operation Operation,
	) (<-chan T, error)
}
