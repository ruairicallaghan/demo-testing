package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockShapeService is a stubbed implementation of the ShapeFactory interface.
type MockShapeService struct{}

// GetShape is a mocked implementation of the GetShape function. This is where we override
// the dependency on the external system and gain control of the output.
func (m *MockShapeService) GetShape(shape string) string {
	return shape
}

func TestGenerateShapes(t *testing.T) {
	sf := &MockShapeService{}
	service := Service{
		ShapeService: sf,
	}

	shapes := service.GenerateShapes(3)
	fmt.Println("Shapes: ", shapes)

	assert.Equal(t, []string{"square", "square", "square"}, shapes)
}
