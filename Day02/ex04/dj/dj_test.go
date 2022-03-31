package dj

import (
	"ex02/product"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// func TestDj(T *testing.T) {
// 	assert := assert.New(T)

// 	gift, _ := product.NewProduct(1500, "ring")
// 	ups, _ := NewDj("John", "Jane", gift)

// 	// assert.Equal(ups.To, "John", "Must Be Equal")
// 	// assert.Equal(ups.From, "Jane", "Must Be Equal")
// }

func TestInfo(T *testing.T) {

	var ups Dj
	assert := assert.New(T)
	assert.Equal(ups.Info(), false, "Must found false when missing field")
	ups.SetTo("Jane")
	assert.Equal(ups.Info(), false, "Must found false when missing field")
	ups.SetFrom("John")
	assert.Equal(ups.Info(), false, "Must found false when missing field")

	val, _ := product.NewProduct(0, "denial")
	ups.SetProduct(val)
	fmt.Println(val)
	assert.Equal(ups.Info(), true, "All fields are completed, need found True")
}
