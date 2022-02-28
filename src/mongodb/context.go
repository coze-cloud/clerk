package mongodb

import (
	"context"
	"time"

	clerk "github.com/coze-hosting/clerk/src"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type monogdbContext struct {
	client *mongo.Client
}

func newMongoContext(client *mongo.Client) *monogdbContext {
	return &monogdbContext{
		client: client,
	}
}

func (c *monogdbContext) Create(collection *clerk.Collection, data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.client.Database(collection.Database).Collection(collection.Name).
		InsertOne(ctx, data)
	return err
}

func (c *monogdbContext) Update(collection *clerk.Collection, filter map[string]interface{}, entity interface{}, upsert bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Replace().SetUpsert(upsert)
	_, err := c.client.Database(collection.Database).Collection(collection.Name).
		ReplaceOne(ctx, filter, entity, opts)
	return err
}

func (c *monogdbContext) Query(collection *clerk.Collection, filter map[string]interface{}) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := c.client.Database(collection.Database).Collection(collection.Name).
		Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	results := []interface{}{}
	for cursor.Next(ctx) {
		var result interface{}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
