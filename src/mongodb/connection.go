package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongodbConnection struct {
	ctx    context.Context
	client *mongo.Client
}

func NewMongoConnection(ctx context.Context, url string) (*MongodbConnection, error) {
	var err error

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, 5*time.Second)
	defer timeoutCancel()
	if err = client.Ping(timeoutCtx, readpref.Primary()); err != nil {
		return nil, err
	}

	return &MongodbConnection{
		ctx:    ctx,
		client: client,
	}, nil
}

func (c *MongodbConnection) Close(handler func(err error)) {
	err := c.client.Disconnect(c.ctx)
	if handler != nil {
		handler(err)
	}
}
