package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs2d(t *testing.T) {
	// create an object used for the test
	square := Shape{
		Name:       "square",
		Dimensions: 2,
	}

	// call the function we want to test
	square2d := is2d(square)

	// ensure the result is as expected
	assert.Equal(t, true, square2d)
}
