package clerk

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestNewMongoUpdateCommand(t *testing.T) {
	// Arrange
	collection := NewDatabase("ExampleDatabase").
		GetCollection("ExampleCollection")

	entity := "Hello World"

	// Act
	command := NewMongoUpdateCommand(collection, entity)

	// Assert
	assert.Equal(t, collection, command.collection)
	assert.Equal(t, entity, command.entity)
}

func TestMongoUpdateCommand_Where(t *testing.T) {
	// Arrange
	collection := NewDatabase("ExampleDatabase").
		GetCollection("ExampleCollection")

	entity := "Hello World"

	filterKey := "_id"
	filterValue := "example"
	filter := bson.E{Key: filterKey, Value: filterValue}

	// Act
	command := NewMongoUpdateCommand(collection, entity).
		Where(filterKey, filterValue)

	// Assert
	assert.Equal(t, filter, command.filter[0])
}

func TestMongoUpdateCommand_WithUpsert(t *testing.T) {
	// Arrange
	collection := NewDatabase("ExampleDatabase").
		GetCollection("ExampleCollection")

	entity := "Hello World"

	// Act
	command := NewMongoUpdateCommand(collection, entity).
		WithUpsert()

	// Assert
	assert.True(t, command.upsert)
}
