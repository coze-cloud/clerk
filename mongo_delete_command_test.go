package clerk

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestNewMongoDeleteCommand(t *testing.T) {
	// Arrange
	collection := NewDatabase("ExampleDatabase").
		GetCollection("ExampleCollection")

	// Act
	command := NewMongoDeleteCommand(collection)

	// Assert
	assert.Equal(t, collection, command.collection)
}

func TestMongoDeleteCommand_Where(t *testing.T) {
	// Arrange
	collection := NewDatabase("ExampleDatabase").
		GetCollection("ExampleCollection")

	filterKey := "_id"
	filterValue := "example"
	filter := bson.E{Key: filterKey, Value: filterValue}

	// Act
	command := NewMongoDeleteCommand(collection).Where(filterKey, filterValue)

	// Assert
	assert.Equal(t, filter, command.filter[0])
}