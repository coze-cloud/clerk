package clerk

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type mongoCommandHandler struct {
	CommandHandler // Interface

	collection *mongo.Collection
}

func newMongoCommandHandler(collection *mongo.Collection) CommandHandler {
	handler := new(mongoCommandHandler)

	handler.collection = collection

	return handler
}

func (handler mongoCommandHandler) Create(entity interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := handler.collection.InsertOne(ctx, entity)
	return err
}

func (handler mongoCommandHandler) Update(filter interface{}, entity interface{}, upsert bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	opts := options.Replace().SetUpsert(upsert)
	_, err := handler.collection.ReplaceOne(ctx, filter, entity, opts)
	return err
}

func (handler mongoCommandHandler) Delete(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := handler.collection.DeleteOne(ctx, filter)
	return err
}

