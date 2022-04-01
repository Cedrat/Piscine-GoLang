package main

import (
	"ex01/fibo"
	"fmt"
)

func main() {
	f := fibo.Fibo()
	fmt.Printf("%T\n", f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(fibo.Fibo())
}
