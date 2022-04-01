package main

import (
	"ex02/fibo"
	"fmt"
)

func main() {
	f, _ := fibo.NewFibo()
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", f.Next(f))
	fmt.Println(f.Next(f))
	fmt.Println(f.Next(f))
	fmt.Println(f.Next(f))
	fmt.Println(f.Next(f))
	fmt.Println(f.Next(f))
	fmt.Println(f.Next(f))
}
