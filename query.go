package clerk

type Query interface {
	handle(handler QueryHandler) (Iterator, error)

	getCollection() Collection
}
