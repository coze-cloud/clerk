package clerk

import "go.mongodb.org/mongo-driver/bson"

type MongoDeleteCommand struct {
	collection Collection
	filter bson.D
}

func NewMongoDeleteCommand(collection Collection) MongoDeleteCommand {
	return MongoDeleteCommand{collection: collection}
}

func (command MongoDeleteCommand) Where(key string, value interface{}) MongoDeleteCommand {
	command.filter = append(command.filter, bson.E{Key: key, Value: value})
	return command
}

func (command MongoDeleteCommand) handle(handler CommandHandler) error {
	return handler.Delete(command.filter)
}

func (command MongoDeleteCommand) getCollection() Collection {
	return command.collection
}