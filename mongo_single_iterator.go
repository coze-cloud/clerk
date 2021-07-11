package clerk

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mcuadros/go-defaults.v1"
)

type mongoSingleIterator struct {
	Iterator // Interface

	hasNext bool `default:"true"`
	result  *mongo.SingleResult
}

func newMongoSingleIterator(result *mongo.SingleResult) Iterator {
	iterator := new(mongoSingleIterator)
	defaults.SetDefaults(iterator)

	iterator.result = result

	return iterator
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
