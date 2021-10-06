package clerk

import "go.mongodb.org/mongo-driver/bson"

type MongoUpdateCommand struct {
	collection Collection
	filter bson.D
	entity interface{}
	upsert bool
}

func NewMongoUpdateCommand(collection Collection, entity interface{}) MongoUpdateCommand {
	return MongoUpdateCommand{
		collection: collection,
		entity: entity,
	}
}

func (command MongoUpdateCommand) Where(key string, value interface{}) MongoUpdateCommand {
	command.filter = append(command.filter, bson.E{Key: key, Value: value})
	return command
}

func (command MongoUpdateCommand) WithUpsert() MongoUpdateCommand {
	command.upsert = true
	return command
}

func (command MongoUpdateCommand) handle(handler CommandHandler) error {
	return handler.Update(command.filter, command.entity, command.upsert)
}

func (command MongoUpdateCommand) getCollection() Collection {
	return command.collection
}
