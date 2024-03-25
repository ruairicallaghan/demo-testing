package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs2d(t *testing.T) {
	square := Shape{
		Name:       "square",
		Dimensions: 2,
	}
	square2d := is2d(square)

	assert.Equal(t, true, square2d)
}
