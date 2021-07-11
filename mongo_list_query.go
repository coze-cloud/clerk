package clerk

type mongoListQuery struct {
	Query // Interface

	collection Collection
}

func NewMongoListQuery(collection Collection) *mongoListQuery {
	query := new(mongoListQuery)

	query.collection = collection

	return query
}

func (query mongoListQuery) GetCollection() Collection {
	return query.collection
}

func (query mongoListQuery) Handle(handler QueryHandler) (Iterator, error) {
	return handler.RetrieveAll()
}