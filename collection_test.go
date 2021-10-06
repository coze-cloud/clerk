package clerk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCollection(t *testing.T) {
	// Arrange
	database := NewDatabase("ExampleDatabase")
	name := "ExampleCollection"

	// Act
	result := NewCollection(database, name)

	// Assert
	assert.Equal(t, database, result.database)
	assert.Equal(t, name, result.name)
}