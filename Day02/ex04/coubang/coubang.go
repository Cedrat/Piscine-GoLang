package coubang

import (
	"errors"
	"ex04/product"
	"fmt"
)

type Coubang struct {
	to         string
	from       string
	membership bool
	product    *product.Product
}

func NewCoubang(to string, from string, membership bool, product *product.Product) (*Coubang, error) {
	var err error

	if product == nil {
		err = errors.New("Zero product")
	}
	a := Coubang{to, from, membership, product}
	return &a, err
}

func (coubang *Coubang) Info() bool {

	switch {
	case len(coubang.to) == 0:
		fmt.Println("ToNotPResent")
		return false
	case len(coubang.from) == 0:
		fmt.Println("FromNotPResent")
		return false
	case coubang.product == nil:
		fmt.Println("ProductNil")
		return false
	case coubang.membership == false:
		fmt.Println("Not a member of coubang")
		return false
	}
	return true
}

func (coubang *Coubang) Send() {
	if coubang.Info() {
		fmt.Println(coubang.to+" just sending to "+coubang.from+" :\n"+coubang.product.Name()+" of a value of", coubang.product.Price())
	} else {
		fmt.Println("Cannot be launch, because he is not a membership or product missing")
	}
}

func (coubang *Coubang) SendRocket() {
	if coubang.Info() {
		fmt.Println(coubang.to+" just sending with a rocket an object to "+coubang.from+" :\n"+coubang.product.Name()+" of a value of", coubang.product.Price())
	} else {
		fmt.Println("Cannot be launch, because he is not a membership or product missing")
	}
}
