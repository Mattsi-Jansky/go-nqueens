package nqueens_test

import (
	"github.com/stretchr/testify/assert"
	"nqueens"
	"testing"
)

func TestSizeFour(t *testing.T) {
	actual := nqueens.Nqueens(4)

	assert.Equal(t, [][]int{
		{1, 3, 0, 2},
		{2, 0, 3, 1},
	}, actual)
}
