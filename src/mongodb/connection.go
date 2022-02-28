package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbConnection struct {
	ctx    context.Context
	client *mongo.Client
}

func NewMongoConnection(url string) (*MongodbConnection, error) {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	return &MongodbConnection{
		ctx:    ctx,
		client: client,
	}, nil
}

func (c *MongodbConnection) Context() *MonogdbContext {
	return newMongoContext(c.client)
}

func (c *MongodbConnection) Close(handler func(err error)) {
	err := c.client.Disconnect(c.ctx)
	if handler != nil {
		handler(err)
	}
}
