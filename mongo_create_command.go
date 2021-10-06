package clerk

type MongoCreateCommand struct {
	collection Collection
	entity interface{}
}

func NewMongoCreateCommand(collection Collection, entity interface{}) MongoCreateCommand {
	return MongoCreateCommand{collection: collection, entity: entity}
}

func (command MongoCreateCommand) handle(handler CommandHandler) error {
	return handler.Create(command.entity)
}

func (command MongoCreateCommand) getCollection() Collection {
	return command.collection
}



