package clerk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMongoCreateCommand(t *testing.T) {
	// Arrange
	collection := NewDatabase("ExampleDatabase").
		GetCollection("ExampleCollection")

	entity := "Hello World"

	// Act
	command := NewMongoCreateCommand(collection, entity)

	// Assert
	assert.Equal(t, collection, command.collection)
	assert.Equal(t, entity, command.entity)
}
