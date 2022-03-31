package product

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct(T *testing.T) {
	assert := assert.New(T)

	a, _ := NewProduct(30000, "car")
	assert.Equal(uint(30000), a.Price(), "Value need to be the same")
	assert.Equal("car", a.name, "Value need to be the same")

}
