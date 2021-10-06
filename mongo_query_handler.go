package clerk

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mongoQueryHandler struct {
	collection *mongo.Collection
}

func newMongoQueryHandler(collection *mongo.Collection) QueryHandler {
	return &mongoQueryHandler{collection: collection}
}

func (handler mongoQueryHandler) RetrieveAll() (Iterator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	cursor, err := handler.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	return newMongoIterator(ctx, cursor), nil
}

func (handler mongoQueryHandler) Retrieve(filter interface{}) (Iterator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	result := handler.collection.FindOne(ctx, filter);

	return newMongoSingleIterator(result), nil
}
