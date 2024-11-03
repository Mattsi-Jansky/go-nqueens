package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSizeFour(t *testing.T) {
	actual := nqueens(4)

	assert.Equal(t, [][]int{
		{1, 3, 0, 2},
		{2, 0, 3, 1},
	}, actual)
}

func nqueens(size int) [][]int {
	return [][]int{
		{},
	}
}
