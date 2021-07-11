package clerk

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type mongoConnection struct {
	Connection // Interface

	ctx context.Context
	client *mongo.Client
}

func NewMongoConnection(connectionString string) (Connection, error) {
	connection := new(mongoConnection)

	var err error

	var cancel context.CancelFunc
	connection.ctx, cancel = context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)
	connection.client, err = mongo.Connect(connection.ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = connection.client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func (connection mongoConnection) SendCommand(command Command) error {
	databaseName := command.GetCollection().Database.Name
	collectionName := command.GetCollection().Name

	collection := connection.client.Database(databaseName).Collection(collectionName)
	commandHandler := newMongoCommandHandler(collection)

	return command.Handle(commandHandler)
}

func (connection mongoConnection) Close(errorHandler func(err error)) {
	err := connection.client.Disconnect(connection.ctx)
	if err != nil && errorHandler != nil {
		errorHandler(err)
	}
}