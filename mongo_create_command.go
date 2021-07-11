package clerk

type mongoCreateCommand struct {
	Command // Interface

	collection Collection
	entity interface{}
}

func NewMongoCreateCommand(collection Collection, entity interface{}) *mongoCreateCommand {
	command := new(mongoCreateCommand)

	command.collection = collection
	command.entity = entity

	return command
}

func (command mongoCreateCommand) GetCollection() Collection {
	return command.collection
}

func (command mongoCreateCommand) Handle(handler CommandHandler) error {
	return handler.Create(command.entity)
}


