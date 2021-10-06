package clerk

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoIterator struct {
	ctx context.Context
	cursor *mongo.Cursor
}

func newMongoIterator(ctx context.Context, cursor *mongo.Cursor) Iterator {
	return &mongoIterator{ctx: ctx, cursor: cursor}
}

func (iterator mongoIterator) Next() bool {
	return iterator.cursor.Next(iterator.ctx)
}

func (iterator mongoIterator) Decode(entity interface{}) error {
	return iterator.cursor.Decode(entity)
}

func (iterator mongoIterator) Close() {
	_ = iterator.cursor.Close(iterator.ctx)
}

func (iterator mongoIterator) Single(entity interface{}) error {
	defer iterator.Close()
	if iterator.Next() {
		return iterator.Decode(entity)
	}
	return nil
}

