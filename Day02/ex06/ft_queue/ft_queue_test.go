package ft_queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// func CheckEqualQueue(a Ft_queue, b Ft_queue) bool {
// if (a.Len() != b.Len()) {
// return false
// }
//
// for i:=0 ; i < a.Len() ; i++ {
// if a
// }
// }

func TestEmpty(T *testing.T) {
	assert := assert.New(T)

	var a Ft_queue

	a = Ft_queue{}

	assert.True(a.Empty(), true, "ft_queue is empty, return value need to be true")
	a.Push(2)
	assert.False(a.Empty(), false, "ft_queue is  not empty, return value need to be false")
}

func TestLen(T *testing.T) {
	assert := assert.New(T)

	var a Ft_queue
	a = Ft_queue{}

	assert.Equal(0, a.Len(), "ft_Queue empty size == 0")

	a.Push(4)
	assert.Equal(1, a.Len(), "ft_Queue size expected == 1")

	a.Push(5)
	a.Push(5)
	assert.Equal(3, a.Len(), "ft_Queue size expected == 3")
}

func TestPop(T *testing.T) {
	assert := assert.New(T)

	a := Ft_queue{}

	a.Push(2)
	a.Push(3)
	a.Push(4)
	b := a.Pop()
	assert.Equal(2, b, "Bad value of first index")
	a.Pop()
	b = a.Pop()
	assert.Nil(b, "zero element in Ft_queue")

	// assert.NotPanics(a.Pop(), "Need to not panic")
}
