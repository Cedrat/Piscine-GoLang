package calc

import "fmt"

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

func (calc *Calculator) Run(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			if b == 0 {
				fmt.Println("Div or Mod by zero")
			}
			if calc.cnt < 0 {
				calc.cnt = 0
			}
		}
	}()

	calc.cnt++
	if calc.cnt < 0 {
		panic("Embedder counter overflow")
	}

	return calc.f(a, b)
}
