package clerk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMongoListQuery(t *testing.T) {
	// Arrange
	collection := NewDatabase("ExampleDatabase").
		GetCollection("ExampleCollection")

	// Act
	command := NewMongoListQuery(collection)

	// Assert
	assert.Equal(t, collection, command.collection)
}
