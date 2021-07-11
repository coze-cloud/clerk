package clerk

type Query interface {
	GetCollection() Collection

	Handle(handler QueryHandler) (Iterator, error)
}
