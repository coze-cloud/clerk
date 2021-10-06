package clerk

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoSingleIterator struct {
	hasNext bool
	result  *mongo.SingleResult
}

func newMongoSingleIterator(result *mongo.SingleResult) Iterator {
	return &mongoSingleIterator{hasNext: true, result: result}
}

func (iterator mongoSingleIterator) Next() bool {
	return iterator.hasNext
}

func (iterator *mongoSingleIterator) Decode(entity interface{}) error {
	iterator.hasNext = false
	return iterator.result.Decode(entity)
}

func (iterator mongoSingleIterator) Close() {
	// Doing nothing here is intended
}

func (iterator mongoSingleIterator) Single(entity interface{}) error {
	return iterator.Decode(entity)
}
