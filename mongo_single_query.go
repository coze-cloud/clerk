package clerk

import "go.mongodb.org/mongo-driver/bson"

type mongoSingleQuery struct {
	Query // Interface

	collection Collection
	filter bson.D
}

func NewMongoSingleQuery(collection Collection) *mongoSingleQuery {
	query := new(mongoSingleQuery)

	query.collection = collection

	return query
}

func (query mongoSingleQuery) Where(key string, value interface{}) mongoSingleQuery {
	query.filter = append(query.filter, bson.E{Key: key, Value: value})
	return query
}

func (query mongoSingleQuery) GetCollection() Collection {
	return query.collection
}

func (query mongoSingleQuery) Handle(handler QueryHandler) (Iterator, error) {
	return handler.Retrieve(query.filter)
}