package product

import "errors"

type Product struct {
	name  string
	price uint
}

func (self Product) Name() string {
	return self.name
}

func (self Product) Price() uint {
	return self.price
}

func NewProduct(price uint, name string) (*Product, error) {

	var product Product

	var err error

	if len(name) == 0 {
		err = errors.New("No Name String")
	}
	product.name = name
	product.price = price

	return &product, err
}
