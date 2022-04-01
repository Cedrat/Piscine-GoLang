package fibo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibo(T *testing.T) {
	assert := assert.New(T)

	// var a int
	a, _ := NewFibo()
	assert.Equal(a.Next(a), 0)
	assert.Equal(a.Next(a), 1)
	assert.Equal(a.Next(a), 1)
	assert.Equal(a.Next(a), 2)
	assert.Equal(a.Next(a), 3)
	assert.Equal(a.Next(a), 5)
	assert.Equal(a.Next(a), 8)
	assert.Equal(a.Next(a), 13)
	assert.Equal(a.Next(a), 21)
	assert.Equal(a.Next(a), 34)
	assert.Equal(a.Next(a), 55)
}
