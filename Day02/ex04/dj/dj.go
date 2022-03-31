package dj

import (
	"ex04/product"
	"fmt"
)

type Dj struct {
	to      string
	from    string
	product *product.Product
}

func (dj *Dj) Info() bool {

	switch {
	case len(dj.to) == 0:
		fmt.Println("ToNotPResent")
		return false
	case len(dj.from) == 0:
		fmt.Println("FromNotPResent")
		return false
	case dj.product == nil:
		fmt.Println("ProductNil")
		return false
	}
	return true
}

func (dj *Dj) Send() {
	fmt.Println(dj.to+" just sending to "+dj.from+" :\n"+dj.product.Name()+" of a value of", dj.product.Price())

}

func (dj *Dj) SetTo(to string) {
	dj.to = to
}

func (dj *Dj) SetFrom(from string) {
	dj.from = from
}

func (dj *Dj) SetProduct(product *product.Product) {
	dj.product = product
}

func NewDj(to string, from string, product *product.Product) (*Dj, error) {
	dj := Dj{to, from, product}

	return &dj, nil

}
