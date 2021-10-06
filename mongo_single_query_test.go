package clerk

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestNewMongoSingleQuery(t *testing.T) {
	// Arrange
	collection := NewDatabase("ExampleDatabase").
		GetCollection("ExampleCollection")

	// Act
	command := NewMongoSingleQuery(collection)

	// Assert
	assert.Equal(t, collection, command.collection)
}

func TestMongoSingleQuery_Where(t *testing.T) {
	// Arrange
	collection := NewDatabase("ExampleDatabase").
		GetCollection("ExampleCollection")

	filterKey := "_id"
	filterValue := "example"
	filter := bson.E{Key: filterKey, Value: filterValue}

	// Act
	command := NewMongoSingleQuery(collection).Where(filterKey, filterValue)

	// Assert
	assert.Equal(t, filter, command.filter[0])
}