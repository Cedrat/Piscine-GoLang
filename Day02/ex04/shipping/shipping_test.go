package shipping

import (
	"ex04/coubang"
	"ex04/dj"
	"ex04/product"
	// "github.com/stretchr/testify/assert"
	"testing"
)

func TestShipping(T *testing.T) {
	// assert := assert.New(T)

	a, _ := product.NewProduct(1500, "ring")
	b, _ := dj.NewDj("Me", "to", a)
	c, _ := coubang.NewCoubang("p1", "p2", true, a)

	shipping(b)
	shipping(c)

}
