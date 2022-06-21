package mongodb

import (
	"context"
	"strings"

	clerk "github.com/coze-cloud/clerk/src"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var operationMap = map[clerk.Operation]string{
	clerk.Create: "insert",
	clerk.Delete: "delete",
	clerk.Update: "replace|update",
}

type MongodbOperator[T any] struct {
	client *mongo.Client
}

func NewMongoOperator[T any](connection *MongodbConnection) *MongodbOperator[T] {
	return &MongodbOperator[T]{
		client: connection.client,
	}
}

func (c *MongodbOperator[T]) Create(
	ctx context.Context,
	collection *clerk.Collection,
	data T,
) error {
	_, err := c.client.
		Database(collection.Database).
		Collection(collection.Name).
		InsertOne(ctx, data)

	return err
}

func (c *MongodbOperator[T]) Update(
	ctx context.Context,
	collection *clerk.Collection,
	filter map[string]any,
	data T,
	upsert bool,
) error {
	opts := options.
		Replace().
		SetUpsert(upsert)

	_, err := c.client.
		Database(collection.Database).
		Collection(collection.Name).
		ReplaceOne(ctx, filter, data, opts)

	return err
}

func (c *MongodbOperator[T]) Delete(
	ctx context.Context,
	collection *clerk.Collection,
	filter map[string]any,
) error {
	_, err := c.client.
		Database(collection.Database).
		Collection(collection.Name).
		DeleteOne(ctx, filter)

	return err
}

func (c *MongodbOperator[T]) Query(
	ctx context.Context,
	collection *clerk.Collection,
	filter map[string]any,
) (<-chan T, error) {
	opts := options.
		Find()

	cursor, err := c.client.
		Database(collection.Database).
		Collection(collection.Name).
		Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	out := make(chan T)

	go func() {
		for cursor.Next(ctx) {
			var data T
			cursor.Decode(&data)
			out <- data
		}
		close(out)
	}()

	return out, nil
}

func (c *MongodbOperator[T]) Watch(
	ctx context.Context,
	collection *clerk.Collection,
	operation clerk.Operation,
) (<-chan T, error) {
	opts := options.
		ChangeStream().
		SetFullDocument(options.UpdateLookup)

	stream, err := c.client.
		Database(collection.Database).
		Collection(collection.Name).
		Watch(ctx, mongo.Pipeline{}, opts)
	if err != nil {
		return nil, err
	}

	out := make(chan T)

	go func() {
		for stream.Next(ctx) {
			var event bson.M
			if err := stream.Decode(&event); err != nil {
				continue
			}

			operationType := event["operationType"].(string)
			if !strings.Contains(operationMap[operation], operationType) {
				continue
			}

			document := event["documentKey"].(bson.M)
			if operationType != "delete" {
				document = event["fullDocument"].(bson.M)
			}

			documentData, err := bson.Marshal(document)
			if err != nil {
				continue
			}

			var data T
			if err := bson.Unmarshal(documentData, &data); err != nil {
				continue
			}

			out <- data
		}
		close(out)
	}()

	return out, nil
}
