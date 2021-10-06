package clerk

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type mongoConnection struct {
	ctx context.Context
	client *mongo.Client
}

func NewMongoConnection(connectionString string) (Connection, error) {
	var err error

	var cancel context.CancelFunc
	connectionCtx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(connectionCtx, clientOptions)
	if err != nil {
		return nil, err
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = client.Ping(pingCtx, nil)
	if err != nil {
		return nil, err
	}

	return &mongoConnection{
		ctx: connectionCtx,
		client: client,
	}, nil
}

func (c mongoConnection) SendQuery(query Query) (Iterator, error) {
	databaseName := query.getCollection().database.name
	collectionName := query.getCollection().name

	collection := c.client.Database(databaseName).Collection(collectionName)
	queryHandler := newMongoQueryHandler(collection)

	return query.handle(queryHandler)
}

func (c mongoConnection) SendCommand(command Command) error {
	databaseName := command.getCollection().database.name
	collectionName := command.getCollection().name

	collection := c.client.Database(databaseName).Collection(collectionName)
	commandHandler := newMongoCommandHandler(collection)

	return command.handle(commandHandler)
}

func (c mongoConnection) Close(handler func(err error)) {
	if err := c.client.Disconnect(c.ctx); handler != nil && err != nil {
		handler(err)
	}
}