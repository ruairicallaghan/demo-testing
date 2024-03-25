package main

import (
	"fmt"
	"time"
)

// Service is an API Service for handling Shapes.
type Service struct {
	ShapeService ShapeFactory
}

// ShapeFactory is an interface for getting Shapes.
type ShapeFactory interface {
	GetShape(string) string
}

// ThirdPartyShapeService is an example of an external system that provides shapes.
// e.g. a GCP service
type ThirdPartyShapeService struct{}

// GetShape is the method attached to ThirdPartyShapeService that interacts with it to
// retrieve shapes.
func (t *ThirdPartyShapeService) GetShape(shape string) string {
	// Complex logic that takes time or has an external dependency
	// Simulated here with a Sleep command.
	time.Sleep(3 * time.Second)

	// Return result of logic
	return shape
}

// GenerateShapes is a method in our service that hanles generating a list of shapes.
// Within this function is a third party call, that we have no control over, e.g.
// its availability, how perfomant it is, etc.
func (s *Service) GenerateShapes(number int) []string {
	var shapes []string
	for i := 1; i <= number; i++ {
		// Here is the external dependency.
		shape := s.ShapeService.GetShape("square")
		shapes = append(shapes, shape)
	}

	return shapes
}

func main() {
	// Define a new ShapeService, composed of our third party service object
	sf := &ThirdPartyShapeService{}
	service := Service{
		ShapeService: sf,
	}

	// Generate some shapes
	shapes := service.GenerateShapes(3)
	fmt.Println("Shapes: ", shapes) // Shapes:  [square square square]
}
