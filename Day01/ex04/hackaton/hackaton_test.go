package hackaton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHackaton(T *testing.T) {
	assert := assert.New(T)

	a := []int{2, 1, 3, 0, 5, 4, 0, 1}
	b := Match(a)
	result := [][]int{{0, 2, 3}, {1}, {4, 5}, {6, 7}}

	assert.Equal(result, b, "The data would be the same")

	a = []int{2, 1, 3, 2, 5, 4, 0, 1}
	b = Match(a)
	result = [][]int{{2, 3}, {1}, {4, 5}, {0, 6, 7}}

	assert.Equal(result, b, "The data would be the same")
}
