package clerk

type MongoListQuery struct {
	collection Collection
}

func NewMongoListQuery(collection Collection) MongoListQuery {
	return MongoListQuery{collection: collection}
}

func (query MongoListQuery) handle(handler QueryHandler) (Iterator, error) {
	return handler.RetrieveAll()
}

func (query MongoListQuery) getCollection() Collection {
	return query.collection
}