package clerk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	// Arrange
	name := "ExampleDatabase"

	// Act
	result := NewDatabase(name)

	// Assert
	assert.Equal(t, name, result.name)
}

func TestDatabase_GetCollection(t *testing.T) {
	// Arrange
	database := NewDatabase("ExampleDatabase")

	collectionName := "ExampleCollection"
	collection := NewCollection(database, collectionName)

	// Act
	result := database.GetCollection(collectionName)

	// Assert
	assert.Equal(t, result, collection)
}
