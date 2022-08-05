package clerk

import "context"

type Operation int

const (
	Create Operation = iota
	Delete
	Update
)

var (
	Any = []Operation{Create, Delete, Update}
)

func (o Operation) String() string {
	switch o {
	case Create:
		return "create"
	case Delete:
		return "delete"
	case Update:
		return "update"
	default:
		return "unknown"
	}
}

type envelope[T any] struct {
	operation Operation
	data      T
}

func (e envelope[T]) Operation() Operation {
	return e.operation
}

func (e envelope[T]) Data() T {
	return e.data
}

type watch[T any] struct {
	collection *Collection
	operations []Operation
}

func NewWatch[T any](collection *Collection, operations ...Operation) *watch[T] {
	return &watch[T]{
		collection: collection,
		operations: operations,
	}
}

func (w *watch[T]) Execute(ctx context.Context, watcher Watcher[T]) (<-chan envelope[T], error) {
	out := make(chan envelope[T])

	for _, operation := range w.operations {
		stream, err := watcher.Watch(ctx, w.collection, operation)
		if err != nil {
			return nil, err
		}

		go func(operation Operation) {
			for data := range stream {
				out <- envelope[T]{
					operation: operation,
					data:      data,
				}
			}
		}(operation)
	}

	return out, nil
}
