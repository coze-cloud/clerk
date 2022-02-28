package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodbConnection struct {
	ctx    context.Context
	client *mongo.Client
}

func NewMongoConnection(url string) (*mongodbConnection, error) {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	return &mongodbConnection{
		ctx:    ctx,
		client: client,
	}, nil
}

func (c *mongodbConnection) Context() *monogdbContext {
	return newMongoContext(c.client)
}

func (c *mongodbConnection) Close(handler func(err error)) {
	err := c.client.Disconnect(c.ctx)
	if handler != nil {
		handler(err)
	}
}
