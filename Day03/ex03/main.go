package main

import (
	"ex03/calc"
	"fmt"
)

func main() {
	c := calc.NewCalculator(func(a, b int) int {
		return a + b
	})
	fmt.Println(c.Run(3, 2))
	fmt.Println(c.Count())
}
