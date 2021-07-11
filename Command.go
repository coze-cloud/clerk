package clerk

type Command interface {
	GetCollection() Collection

	Handle(handler CommandHandler) error
}