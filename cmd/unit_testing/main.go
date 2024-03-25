package main

import "fmt"

type Shape struct {
	Name       string
	Dimensions int
}

// is2d checks if a given shape has 2 Dimensions.
func is2d(shape Shape) bool {
	if shape.Dimensions == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	square := Shape{
		Name:       "square",
		Dimensions: 2,
	}
	square2d := is2d(square)

	cube := Shape{
		Name:       "cube",
		Dimensions: 3,
	}
	cube2d := is2d(cube)

	fmt.Println("square is 2d: ", square2d)
	fmt.Println("cube is 2d: ", cube2d)
}
