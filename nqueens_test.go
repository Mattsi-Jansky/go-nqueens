package nqueens_test

import (
	"github.com/stretchr/testify/assert"
	"nqueens"
	"testing"
)

func TestInvalidState(t *testing.T) {
	actual := nqueens.InvalidState([]int{0, 0})

	assert.Equal(t, true, actual)
}

func TestInvalidStateSomeDiagonal(t *testing.T) {
	actual := nqueens.InvalidState([]int{0, 1})

	assert.Equal(t, true, actual)
}

func TestInvalidStateSomeDiagonal2(t *testing.T) {
	actual := nqueens.InvalidState([]int{0, 2})

	assert.Equal(t, false, actual)
}

func TestSizeFour(t *testing.T) {
	actual := nqueens.Nqueens(4)

	assert.Equal(t, [][]int{
		{1, 3, 0, 2},
		{2, 0, 3, 1},
	}, actual)
}

func TestSizeFive(t *testing.T) {
	actual := nqueens.Nqueens(5)

	assert.Equal(t, 10, len(actual))
}

func TestSizeNine(t *testing.T) {
	actual := nqueens.Nqueens(9)

	assert.Equal(t, 352, len(actual))
}

func BenchmarkSpeedSize(b *testing.B) {
	nqueens.Nqueens(10)
}
