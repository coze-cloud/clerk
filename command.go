package clerk

type Command interface {
	handle(handler CommandHandler) error

	getCollection() Collection
}