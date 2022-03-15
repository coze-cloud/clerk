package mongodb

import (
	"context"
	"time"

	clerk "github.com/coze-cloud/clerk/src"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MonogdbContext struct {
	client *mongo.Client
}

func newMongoContext(client *mongo.Client) *MonogdbContext {
	return &MonogdbContext{
		client: client,
	}
}

func (c *MonogdbContext) Create(collection *clerk.Collection, data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.client.Database(collection.Database).Collection(collection.Name).
		InsertOne(ctx, data)
	return err
}

func (c *MonogdbContext) Update(collection *clerk.Collection, filter map[string]interface{}, data interface{}, upsert bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Replace().SetUpsert(upsert)
	_, err := c.client.Database(collection.Database).Collection(collection.Name).
		ReplaceOne(ctx, filter, data, opts)
	return err
}

func (c *MonogdbContext) Delete(collection *clerk.Collection, filter map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.client.Database(collection.Database).Collection(collection.Name).
		DeleteOne(ctx, filter)
	return err
}

func (c *MonogdbContext) Query(collection *clerk.Collection, filter map[string]interface{}, results interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := c.client.Database(collection.Database).Collection(collection.Name).
		Find(ctx, filter)
	if err != nil {
		return err
	}

	if err := cursor.All(ctx, results); err != nil {
		return err
	}
	return nil
}
