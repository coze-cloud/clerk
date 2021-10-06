package clerk

import "go.mongodb.org/mongo-driver/bson"

type MongoSingleQuery struct {
	collection Collection
	filter bson.D
}

func NewMongoSingleQuery(collection Collection) MongoSingleQuery {
	return MongoSingleQuery{collection: collection}
}

func (query MongoSingleQuery) Where(key string, value interface{}) MongoSingleQuery {
	query.filter = append(query.filter, bson.E{Key: key, Value: value})
	return query
}

func (query MongoSingleQuery) handle(handler QueryHandler) (Iterator, error) {
	return handler.Retrieve(query.filter)
}

func (query MongoSingleQuery) getCollection() Collection {
	return query.collection
}