package clerk

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type mongoCommandHandler struct {
	collection *mongo.Collection
}

func newMongoCommandHandler(collection *mongo.Collection) CommandHandler {
	return &mongoCommandHandler{collection: collection}
}

func (h mongoCommandHandler) Create(entity interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := h.collection.InsertOne(ctx, entity)
	return err
}

func (h mongoCommandHandler) Update(filter interface{}, entity interface{}, upsert bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	opts := options.Replace().SetUpsert(upsert)
	_, err := h.collection.ReplaceOne(ctx, filter, entity, opts)
	return err
}

func (h mongoCommandHandler) Delete(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := h.collection.DeleteOne(ctx, filter)
	return err
}

