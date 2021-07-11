package clerk

import "go.mongodb.org/mongo-driver/bson"

type mongoDeleteCommand struct {
	Command // Interface

	collection Collection
	filter bson.D
}

func NewMongoDeleteCommand(collection Collection) *mongoDeleteCommand {
	command := new(mongoDeleteCommand)

	command.collection = collection

	return command
}

func (command mongoDeleteCommand) Where(key string, value interface{}) mongoDeleteCommand {
	command.filter = append(command.filter, bson.E{Key: key, Value: value})
	return command
}

func (command mongoDeleteCommand) GetCollection() Collection {
	return command.collection
}

func (command mongoDeleteCommand) Handle(handler CommandHandler) error {
	return handler.Delete(command.filter)
}