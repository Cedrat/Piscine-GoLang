package fibo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibo(T *testing.T) {
	assert := assert.New(T)

	// var a int
	a := Fibo()
	assert.Equal(a(), 0)
	assert.Equal(a(), 1)
	assert.Equal(a(), 1)
	assert.Equal(a(), 2)
	assert.Equal(a(), 3)
	assert.Equal(a(), 5)
	assert.Equal(a(), 8)
	assert.Equal(a(), 13)
	assert.Equal(a(), 21)
	assert.Equal(a(), 34)
	assert.Equal(a(), 55)
}
