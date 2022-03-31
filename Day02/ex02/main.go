package main

import (
	"ex02/dj"
	"ex02/product"
)

func main() {
	var a *product.Product
	a, _ = product.NewProduct(1500, "water")

	var b *dj.Dj
	b, _ = dj.NewDj("Expeditor", "receiver", a)
	b.Send()
}
