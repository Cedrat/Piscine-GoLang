package coubang

import (
	"ex03/product"
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoubang(T *testing.T) {
	assert := assert.New(T)

	var ups *Coubang
	var gift *product.Product

	gift = nil

	ups, _ = NewCoubang("John", "Jane", true, gift)
	assert.Equal(false, ups.Info(), "Need to be false, no gift")

	ups.Send()
	gift, _ = product.NewProduct(30, "Happy Meal")
	ups, _ = NewCoubang("John", "Jane", true, gift)
	assert.Equal(true, ups.Info(), "Need to be true, all is right")
	ups.SendRocket()

}
