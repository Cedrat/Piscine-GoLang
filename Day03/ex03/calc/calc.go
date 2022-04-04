package calc

import "fmt"

const MaxInt = int(^uint(0) >> 1)

type Calculator struct {
	f   func(a, b int) int
	cnt int8
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Div(a, b int) int {
	return a / b
}

func NewCalculator(f func(a, b int) int) *Calculator {
	var calc Calculator

	calc.cnt = 0
	calc.f = f

	return &calc
}

func (calc *Calculator) ChangeOperation(f func(a, b int) int) {
	calc.f = f
}
func (calc *Calculator) Count() int8 {
	return calc.cnt
}

func (calc *Calculator) Run(a, b int) (ret int) {

	defer func() {
		if err := recover(); err != nil {
			if b == 0 {
				fmt.Println("Div or Mod by zero")
				ret = MaxInt
			}
			if calc.cnt < 0 {
				calc.cnt = 0
			}
		}
	}()
	calc.cnt++
	ret = calc.f(a, b)

	if calc.cnt < 0 {
		panic("Embedded counter overflow")
	}

	return ret
}
