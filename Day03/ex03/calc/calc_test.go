package calc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculator(T *testing.T) {
	assert := assert.New(T)

	c := NewCalculator(func(a, b int) int {
		return a + b
	})
	assert.Equal(5, c.Run(2, 3), " 2 + 3 = ?")
	assert.Equal(int8(1), c.Count(), "Calc used one time")

	c.ChangeOperation(Div)
	assert.Equal(2, c.Run(10, 5), " 10  / 5 = ?")
}

func TestDivBy0(T *testing.T) {
	assert := assert.New(T)

	c := NewCalculator(Div)

	for i := 1; i < 500; i++ {
		c.Run(10, 5)
		assert.Equal(int8(i%128), c.Count())
	}

	assert.Equal(MaxInt, c.Run(10, 0))
}
