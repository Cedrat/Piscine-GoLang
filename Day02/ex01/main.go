package main

import (
	"ex01/product"
	"fmt"
)

func main() {
	a, _ := product.NewProduct(30000, "car")

	fmt.Println(a.Name())
}
